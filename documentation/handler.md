# **Handler**
*This package is the one responsible of the calls done to the api and also transforming this information for the project specific purposes*

## **FUNCTIONS**

`func Advantage(w http.ResponseWriter, r *http.Request)`
> Advantage is a function which returns an struct of type advantage, this in order to 
> inform the user which pokemon has and advantage over the other. This is achieved by 
> taking into account the types of the first pokemon the second pokemon (we assume here 
> that the pokemons only have one attack from each of their types). Therefore, the 
> advantage can be calculated by simply running an iterative loop over each of the 
> pokemons types calculate the damage that they would do and recieve. Once the DmgTaken 
> and DmgDone variables are calculated, the difference between these two is also 
> calculated substracting the DmgDone from the DmgTaken:
>    - If it is positive this means that the first pokemon has an advantage over the second one
>    - If it is negative this means that the second pokemon has an advantage over the first one
>    - If it is zero this means that neither of the pokemons have an advantage over the other
> This reasoning of why it has the advantage or not its given to the user via the Explanation 
> variable. All of this is returned as a JSON with the help of the parser package

`func CommonMoves(w http.ResponseWriter, r *http.Request)`
> CommonMoves is a function that given a set of pokemons of two or more, a
> languague and a limit to the maximum common moves returned. It returns a
> JSON with the language, the pokemons that it compared and the moves it found
> that they have in common (this moves limited by the maximum allowed by the
> user, with default value of 1)

`func dmgDone(attack []parser.Drelations, defense parser.Pokemon) float32`
> dmgDone is a function which recieves:
>   -An array of the damage relations that the attacking pokemon has, in this case the attacking pokemon is the first pokemon
>   -The defending pokemon, in this case the defending pokemon is the second one

> The function then proceeds to check the damage relations that the attacking
> pokemon has with the defending pokemon types. In order to calculate the
> damage done the function uses a damage variable that is multiplied by two or
> divided by two depending on the damage relation of the attacking pokemon
> whit the defending pokemon. (The damage returned is 0 if a relation of no
> damage to is encountered).

`func dmgTaken(attack parser.Pokemon, defense []parser.Drelations) float32`
> dmgTaken is a function which recieves:

>     -The attacking pokemon, in this case the attacking pokemon is the first one
>     -An array of the damage relations that the defending pokemon has, in this case the defending pokemon is the second pokemon

> The function then proceeds to check the damage relations that the defending
> pokemon has with the attacking pokemon types. In order to calculate the
> damage done the function uses a damage variable that is multiplied by two or
> divided by two depending on the damage relation of the defending pokemon
> whit the attacking pokemon. (The damage returned is 0 if a relation of no
> damage to is encountered).

`func hasMoves(m string, moves []parser.Move) bool`
> hasMove returns true if a move is a move is on given list Inputs:
>   -m:     Move to be compared to agaisnt the list of moves
>   -moves: List of moves in which we are looking for the desired move
> Returns:
>  -true or false depending on if the move was on the list or not

`func swap(i1 int, i2 int, p *[]parser.Pokemon)`
> swap is a function that swaps two given elements of a list.

`func translateMove(lan string, move parser.Node) parser.Node`
> translate translates a move to the desired language Inputs:
>   -lan:   Code of the language to translate to
>   -move:  Move that is going to be translated
> Returns:
>   -move: Translated move

`func typeInDamage(a string, list []parser.Node) bool`
> typeInDamage returns true if a string given to a is found on the given list

`func getCommonMoves(commonM commonMoves, lan string) commonMoves`
> getCommonMoves will return the moves that the given pokemons have in common
> via the commonMoves struct

## **TYPES**

    type advantage struct {
            HasAdvantage bool    `json:"has_advantage"`
            DmgTaken     float32 `json:"damage_taken"`
            DmgDone      float32 `json:"damage_done"`
            Poke1        string  `json:"pokemon_1"`
            Poke2        string  `json:"pokemon_2"`
            Explanation  string  `json:"explanation"`
    }
> advantage is a struct that stores the data of 2 given pokemons telling which one has an advantage over the other with the following data:
>    - HasAdvantage: Does the first pokemon has an advantage over the second one? (true or false)
>    - DmgTaken:             Multiplier of damage recived by the first pokemon from the second pokemon
>    - DmgDone:              Multiplier of damage done by the first pokemon to the second pokemon
>    - Poke1:                Name of the first pokemon
>    - Poke2:                Name of the second pokemon
>    - Explanation:  Explanation on why the first pokemon has an advantage or not over the second pokemon

    type commonMoves struct {
            Language      string        `json:"language"`
            Pokemons      []string      `json:"pokemon"`
            MovesInCommon []parser.Node `json:"moves_in_common"`
    }
> commonMoves is a struct that stores the data of all given pokemons telling
> which moves are the ones they have in common:
>        - Language:                     Code of the language in which the data is going to be stored
>        - Pokemons:                     Data of the pokemon which moves were compared
>        - MovesInCommon:        A list of all the moves they have in common