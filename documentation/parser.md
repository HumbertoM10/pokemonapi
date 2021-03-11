# **Parser**
*This package helps to parse all the requests beign made to the poke api*

## **FUNCTIONS**

`func GetLanguage(res interface{}) string`
> GetLanguage verifies if language exists and returns the name of the
> language. In case the language does not exist, it will return an empty
> string.

`func SendJSON(w http.ResponseWriter, status int, body interface{})`
> SendJSON is a function that is responsabile for sending (posting) the JSON generated

`func writeHeader(w *http.ResponseWriter, status int)`
> writeHeader writes the headers required for setting the application to type
> json and allowing cors

`func GetDamageRelations(url string) Drelations`
> GetDamageRelations gets the damage relations from the passed url and returns it in a Drelations struct

`func GetMove(url string, lan string) Node`
> GetMove gets a move with the provided url and translates it to the language given in lan

`func GetPokemon(pokeNames []string) []Pokemon`
> GetPokemon given an array of two pokemon, will return an array of type Pokemon

TYPES

    type Drelations struct {
            Name      string    `json:"name"`
            Drelation drelation `json:"damage_relations"`
    }
> Drelations stores a type, and the damage relations of said type

    type Move struct {
            Move Node `json:"move"`
    }
> Move stores an Elem containing a move's data

    type Node struct {
            Name string `json:"name"`
            URL  string `json:"url"`
    }
> Node stores a name and a url, is used to store the parts of the information provided by the consummed api

    type Pokemon struct {
            Name  string     `json:"name"`
            Types []pokeType `json:"types"`
            Moves []Move     `json:"moves"`
    }
> Pokemon stores a pokemon's name, types, and moves

    type TranslatedMove struct {
            Name  string `json:"name"`
            Names []lan  `json:"names"`
    }
> TranslatedMove saves the base name of a move in english and its translations
> using the struct lan

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

    type lan struct {
            Name     string `json:"name"`
            Language Node   `json:"language"`
    }
> language stores in Name the name of something in a given language, and
> Language stores the name, and the corresponding node of the Language beign
> used

    type pokeType struct {
            PokeType Node `json:"type"`
    }
> poketype stores a the pokemontype name and its url