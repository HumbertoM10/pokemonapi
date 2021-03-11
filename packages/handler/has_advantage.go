// This package is the one responsible of the calls done to the api and
// also transforming this information for the project specific purposes
package handler

import (
	"net/http"
	"strings"

	"github.com/HumbertoM10/pokemonapi/packages/parser"
)

// advantage is a struct that stores the data of 2 given pokemons telling
// which one has an advantage over the other with the following data:
//	- HasAdvantage:	Does the first pokemon has an advantage over the second one? (true or false)
//	- DmgTaken:		Multiplier of damage recived by the first pokemon from the second pokemon
//	- DmgDone:		Multiplier of damage done by the first pokemon to the second pokemon
//	- Poke1:		Name of the first pokemon
//	- Poke2:		Name of the second pokemon
//	- Explanation:	Explanation on why the first pokemon has an advantage or not over the second pokemon
type advantage struct {
	HasAdvantage bool    `json:"has_advantage"`
	DmgTaken     float32 `json:"damage_taken"`
	DmgDone      float32 `json:"damage_done"`
	Poke1        string  `json:"pokemon_1"`
	Poke2        string  `json:"pokemon_2"`
	Explanation  string  `json:"explanation"`
}

// Advantage is a function which returns an struct of type advantage, this in order to
// inform the user which pokemon has and advantage over the other.
// This is achieved by taking into account the types of the first pokemon the second
// pokemon (we assume here that the pokemons only have one attack from each of their
// types). Therefore, the advantage can be calculated by simply running an iterative
// loop over each of the pokemons types calculate the damage that they would do and
// recieve.
// Once the DmgTaken and DmgDone variables are calculated, the difference between
// these two is also calculated substracting the DmgDone from the DmgTaken:
//	- If it is positive this means that the first pokemon has an advantage over the second one
//	- If it is negative this means that the second pokemon has an advantage over the first one
//	- If it is zero this means that neither of the pokemons have an advantage over the other
// This reasoning of why it has the advantage or not its given to the user via the
// Explanation variable.
// All of this is returned as a JSON with the help of the parser package
func Advantage(w http.ResponseWriter, r *http.Request) {
	const pokeKey = "pokemon"

	pokeNames := strings.Split(r.URL.Query().Get(pokeKey), ",")
	pokeArr := parser.GetPokemon(pokeNames)

	adv := new(advantage)

	if len(pokeArr) >= 2 {
		adv.Poke1 = pokeArr[0].Name
		adv.Poke2 = pokeArr[1].Name
		dmgRelations := []parser.Drelations{}

		for _, t := range pokeArr[0].Types {
			relation := parser.GetDamageRelations(t.PokeType.URL)
			dmgRelations = append(dmgRelations, relation)
		}

		adv.DmgDone = dmgDone(dmgRelations, pokeArr[1])
		adv.DmgTaken = dmgTaken(pokeArr[1], dmgRelations)

		if adv.DmgDone-adv.DmgTaken > 0 {
			adv.HasAdvantage = true
			adv.Explanation = adv.Poke1 + " has an advantage over " + adv.Poke2
		} else if adv.DmgDone-adv.DmgTaken < 0 {
			adv.Explanation = adv.Poke1 + " doesn't have an advantage over " + adv.Poke2 +
				" in fact " + adv.Poke2 + " has an advantage over " + adv.Poke1
		} else {
			adv.Explanation = adv.Poke1 + " doesn't have an advantage over " + adv.Poke2 +
				" nor " + adv.Poke2 + " has an advantage over " + adv.Poke1
		}
	} else {
		adv.Explanation = "An error has ocurred, please make sure you entered two valid pokemons"
	}

	parser.SendJSON(w, http.StatusOK, adv)
}

// dmgDone is a function which recieves:
// 	-An array of the damage relations that the attacking pokemon has, in this case the attacking pokemon is the first pokemon
// 	-The defending pokemon, in this case the defending pokemon is the second one
// The function then proceeds to check the damage relations that the attacking pokemon has with the defending pokemon types.
// In order to calculate the damage done the function uses a damage variable that is multiplied by two or divided by two
// depending on the damage relation of the attacking pokemon whit the defending pokemon. (The damage returned is 0 if a
// relation of no damage to is encountered).
func dmgDone(attack []parser.Drelations, defense parser.Pokemon) float32 {
	var damage float32 = 1.0

	for _, pt1 := range attack {
		for _, pt2 := range defense.Types {
			if typeInDamage(pt2.PokeType.Name, pt1.Drelation.DoubleDmgTo) {
				damage *= 2
				continue
			} else if typeInDamage(pt2.PokeType.Name, pt1.Drelation.HalfDmgTo) {
				damage /= 2
				continue
			} else if typeInDamage(pt2.PokeType.Name, pt1.Drelation.NoDmgTo) {
				return 0.0
			}
		}
	}

	return damage
}

// dmgTaken is a function which recieves:
// 	-The attacking pokemon, in this case the attacking pokemon is the first one
//	-An array of the damage relations that the defending pokemon has, in this case the defending pokemon is the second pokemon
// The function then proceeds to check the damage relations that the defending pokemon has with the attacking pokemon types.
// In order to calculate the damage done the function uses a damage variable that is multiplied by two or divided by two
// depending on the damage relation of the defending pokemon whit the attacking pokemon. (The damage returned is 0 if a
// relation of no damage to is encountered).
func dmgTaken(attack parser.Pokemon, defense []parser.Drelations) float32 {
	var damage float32 = 1.0

	for _, pt1 := range defense {
		for _, pt2 := range attack.Types {
			if damage > 0 {
				if typeInDamage(pt2.PokeType.Name, pt1.Drelation.DoubleDmgFrom) {
					damage *= 2
					continue
				} else if typeInDamage(pt2.PokeType.Name, pt1.Drelation.HalfDmgFrom) {
					damage *= 0.5
					continue
				} else if typeInDamage(pt2.PokeType.Name, pt1.Drelation.NoDmgFrom) {
					damage *= 0
				}
			}
		}
	}

	return damage
}

// typeInDamage returns true if a string given to a is found on the given list
func typeInDamage(a string, list []parser.Node) bool {
	for _, b := range list {
		if b.Name == a {
			return true
		}
	}
	return false
}
