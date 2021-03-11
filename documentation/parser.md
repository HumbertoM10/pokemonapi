# **Parser**
*This package helps to parse all the requests beign made to the poke api*

## **FUNCTIONS**

```go
func GetLanguage(res interface{}) string
```
> GetLanguage verifies if language exists and returns the name of the
> language. In case the language does not exist, it will return an empty
> string.

```go
func SendJSON(w http.ResponseWriter, status int, body interface{})
```
> SendJSON is a function that is responsabile for sending (posting) the JSON generated

```go
func writeHeader(w *http.ResponseWriter, status int)
```
> writeHeader writes the headers required for setting the application to type
> json and allowing cors

```go
func GetDamageRelations(url string) Drelations
```
> GetDamageRelations gets the damage relations from the passed url and returns it in a Drelations struct

```go
func GetMove(url string, lan string) Node
```
> GetMove gets a move with the provided url and translates it to the language given in lan

```go
func GetPokemon(pokeNames []string) []Pokemon
```
> GetPokemon given an array of two pokemon, will return an array of type Pokemon

## **TYPES**

    type DmgRelations struct {
            Name            string          `json:"name"`
            TypeRelation    TypeRelation    `json:"damage_relations"`
    }
> DmgRelations stores a type, and the damage relations of said type

Name | Description
--- | --- 
*Name* | `Name of the type`
*TypeRelation* | `Type of relation of damage that it has`

    type Move struct {
            Move Node `json:"move"`
    }
> Move stores a Node containing a move's data

Name | Description
--- | --- 
*Move* | `Move data`

    type Node struct {
            Name string `json:"name"`
            URL  string `json:"url"`
    }
> Node stores a name and a url, is used to store the parts of the information provided by the consummed api

Name | Description
--- | --- 
*Name* | `Name of the node`
*URL* | `URL of the node`

    type Pokemon struct {
            Name  string     `json:"name"`
            Types []pokeType `json:"types"`
            Moves []Move     `json:"moves"`
    }
> Pokemon stores a pokemon's name, types, and moves

Name | Description
--- | --- 
*Name* | `Name of the pokemon`
*Types* | `Types of the pokemon`
*Moves* | `Moveset of the pokemon`

    type TranslatedMove struct {
            Name  string `json:"name"`
            Tnames []lan  `json:"names"`
    }
> TranslatedMove saves the base name of a move in english and its translations
> using the struct lan

Name | Description
--- | --- 
*Name* | `Base name of the move`
*Tnames* | `Translated names of the move`

    type drelation struct {
            DoubleDmgFrom []Node `json:"double_damage_from"`
            DoubleDmgTo   []Node `json:"double_damage_to"`
            HalfDmgFrom   []Node `json:"half_damage_from"`
            HalfDmgTo     []Node `json:"half_damage_to"`
            NoDmgFrom     []Node `json:"no_damage_from"`
            NoDmgTo       []Node `json:"no_damage_to"`
    }
> drelation is a structure that stores all the damages made from, and two the
> type in question

Name | Description
--- | --- 
*DoubleDmgFrom* | `Types from which this particular type recieves double damage`
*DoubleDmgTo* | `Types from which this particular type does double damage`
*HalfDmgFrom* | `Types from which this particular type recieves halfed damage`
*HalfDmgTo* | `Types from which this particular type does halfed damage`
*NoDmgFrom* | `Types from which this particular type recieves no damage`
*NoDmgTo* | `Types from which this particular type does no damage`

    type lan struct {
            Name     string `json:"name"`
            Language Node   `json:"language"`
    }
> lan stores in Name the name of the move in a given language, and in
> Language it stores the name, and the corresponding node of the Language beign used

Name | Description
--- | --- 
*Name* | `Code of the language eg: en, es, ja`
*Language* | `Translated name of the move to that specific language`

    type pokeType struct {
            PokeType Node `json:"type"`
    }
> poketype stores a the pokemontype name and its url

Name | Description
--- | --- 
*PokeType* | `Pokemon type data`