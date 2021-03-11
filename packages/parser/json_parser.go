//This package helps to parse all the requests beign made to the poke api
package parser

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Pokemon stores a pokemon's name, types, and moves
type Pokemon struct {
	Name  string     `json:"name"`
	Types []PokeType `json:"types"`
	Moves []Move     `json:"moves"`
}

// poketype stores a the pokemontype name and its url
type PokeType struct {
	PokeType Node `json:"type"`
}

// Move stores a Node containing a move's data
type Move struct {
	Move Node `json:"move"`
}

// Node stores a name and a url, is used to store the parts
// of the information provided by the consummed api
type Node struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// GetPokemon given an array of two pokemon, will return an array of type Pokemon
func GetPokemon(pokeNames []string) []Pokemon {
	url := "https://pokeapi.co/api/v2/pokemon/"
	pokeArr := []Pokemon{}

	// Calling the api to get the pokemon
	for _, p := range pokeNames {
		requestURL := url + p
		response, err := http.Get(requestURL)
		if err != nil {
			log.Fatal(err)
		}
		pokemon := &Pokemon{}
		json.NewDecoder(response.Body).Decode(pokemon)
		pokeArr = append(pokeArr, *pokemon)
	}

	return pokeArr
}

// DmgRelations stores a type, and the damage relations of said type
type DmgRelations struct {
	Name         string       `json:"name"`
	TypeRelation TypeRelation `json:"damage_relations"`
}

// TypeRelation is a structure that stores all the types of damage relations that a pokemon has
// (has the same form of the JSON recieved from the pokeapi)
type TypeRelation struct {
	DoubleDmgFrom []Node `json:"double_damage_from"`
	DoubleDmgTo   []Node `json:"double_damage_to"`
	HalfDmgFrom   []Node `json:"half_damage_from"`
	HalfDmgTo     []Node `json:"half_damage_to"`
	NoDmgFrom     []Node `json:"no_damage_from"`
	NoDmgTo       []Node `json:"no_damage_to"`
}

// GetDamageRelations gets the damage relations from the passed url that has the form of:
// https://pokeapi.co/api/v2/type/{name}/ (where name is the name of the pokemon type) and
// returns it in a DmgRelations struct
func GetDamageRelations(url string) DmgRelations {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	relation := &DmgRelations{}
	json.NewDecoder(response.Body).Decode(relation)

	return *relation
}

// GetLanguage verifies if language exists and returns the name of the language.
// In case the language does not exist, it will return an empty string.
func GetLanguage(res interface{}) string {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/language/%v", res)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	lan := &Node{}
	json.NewDecoder(response.Body).Decode(lan)

	if lan.Name == "" {
		lan.Name = "en"
	}

	return lan.Name

}

// TranslatedMove saves the base name of a move in english and its translations using the struct lan
type TranslatedMove struct {
	Name   string `json:"name"`
	Tnames []lan  `json:"names"`
}

// lan stores in Name the name of something in a given language, and Language stores the name, and
// the corresponding node of the Language beign used
type lan struct {
	Name     string `json:"name"`
	Language Node   `json:"language"`
}

// GetMove gets a move with the provided url and translates it to the language given in lan
func GetMove(url string, lan string) Node {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	move := &TranslatedMove{}
	json.NewDecoder(response.Body).Decode(move)

	for _, l := range move.Tnames {
		if l.Language.Name == lan {
			return Node{l.Name, url}
		}
	}
	return Node{}
}
