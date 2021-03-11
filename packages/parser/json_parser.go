//This package helps to parse all the requests beign made to the poke api and also
package parser

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Pokemon stores a pokemon's name, type, and moveset
type Pokemon struct {
	Name   string  `json:"name"`
	Ptypes []ptype `json:"types"`
	Moves  []Move  `json:"moves"`
}

// ptype stores a types name as well as its url
type ptype struct {
	Ptype Elem `json:"type"`
}

// Move stores an Elem containing a move's data
type Move struct {
	Move Elem `json:"move"`
}

// Elem stores a name as well as a url
type Elem struct {
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
		poke := &Pokemon{}
		json.NewDecoder(response.Body).Decode(poke)
		pokeArr = append(pokeArr, *poke)
	}

	return pokeArr
}

// Drelations stores a type, and the damage relations of said type
type Drelations struct {
	Name      string    `json:"name"`
	Drelation drelation `json:"damage_relations"`
}

// drelation is a structure that stores all the damages made from, and two the type in question
type drelation struct {
	DoubleDmgFrom []Elem `json:"double_damage_from"`
	DoubleDmgTo   []Elem `json:"double_damage_to"`
	HalfDmgFrom   []Elem `json:"half_damage_from"`
	HalfDmgTo     []Elem `json:"half_damage_to"`
	NoDmgFrom     []Elem `json:"no_damage_from"`
	NoDmgTo       []Elem `json:"no_damage_to"`
}

// GetDamageRelations gets da damage relations from the passed url and returns it in a Drelations
// struct
func GetDamageRelations(url string) Drelations {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	relation := &Drelations{}
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

	lan := &Elem{}
	json.NewDecoder(response.Body).Decode(lan)

	if lan.Name == "" {
		lan.Name = "en"
	}

	return lan.Name

}

// CompleteMove saves the name, and all the names in different locales of a move
type CompleteMove struct {
	Name  string     `json:"name"`
	Names []language `json:"names"`
}

// language stores in Name the name of a move in a given language, and Language stores the name, and
// url of the language being used
type language struct {
	Name     string `json:"name"`
	Language Elem   `json:"language"`
}

// GetMove gets a move with the provided url and translates it to the language given in lan
func GetMove(url string, lan string) Elem {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	move := &CompleteMove{}
	json.NewDecoder(response.Body).Decode(move)

	for _, l := range move.Names {
		if l.Language.Name == lan {
			return Elem{l.Name, url}
		}
	}
	return Elem{}
}
