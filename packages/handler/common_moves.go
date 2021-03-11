// This package is the one responsible of the calls done to the api and
// also transforming this information for the project specific purposes
package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/HumbertoM10/pokemonapi/packages/parser"
)

// commonMoves is a struct that stores the data of all given pokemons telling
// which moves are the ones they have in common:
//	- Language:			Code of the language in which the data is going to be stored
//	- Pokemons:			Data of the pokemon which moves were compared
//	- MovesInCommon:	A list of all the moves they have in common
type commonMoves struct {
	Language      string        `json:"language"`
	Pokemons      []string      `json:"pokemon"`
	MovesInCommon []parser.Node `json:"moves_in_common"`
}

// CommonMoves is a function that given a set of pokemons of two or more, a languague
// and a limit to the maximum common moves returned. It returns a JSON with the
// language, the pokemons that it compared and the moves it found that they have in
// common (this moves limited by the maximum allowed by the user, with default value
// of 1)
func CommonMoves(w http.ResponseWriter, r *http.Request) {
	const (
		pokeKey  = "pokemon"
		lanKey   = "language"
		limitKey = "limit"
	)

	language := r.URL.Query().Get(lanKey)

	common := commonMoves{
		Pokemons: strings.Split(r.URL.Query().Get(pokeKey), ","),
	}

	common = getCommonMoves(common, language)

	limit, _ := strconv.Atoi(r.URL.Query().Get(limitKey))

	if limit == 0 {
		limit = 1
	}

	if len(common.MovesInCommon) > limit {
		common.MovesInCommon = common.MovesInCommon[:limit]
	}

	parser.SendJSON(w, http.StatusOK, common)

}

// getCommonMoves will return the moves that the given pokemons have in common
// via the commonMoves struct
func getCommonMoves(commonM commonMoves, lan string) commonMoves {
	pokeArr := parser.GetPokemon(commonM.Pokemons)
	commonM.Language = parser.GetLanguage(lan)

	movesInCommon := []parser.Node{}

	smallest := 0
	for i := 0; i < len(pokeArr); i++ {
		if len(pokeArr[i].Moves) < len(pokeArr[smallest].Moves) {
			smallest = i
		}
	}
	swap(0, smallest, &pokeArr)

	for _, m := range pokeArr[0].Moves {
		common := true
		for i := 1; i < len(pokeArr); i++ {
			if hasMoves(m.Move.Name, pokeArr[i].Moves) != true {
				common = false
				break
			}
		}
		if common {
			move := translateMove(commonM.Language, m.Move)
			movesInCommon = append(movesInCommon, move)
		}
	}

	commonM.MovesInCommon = movesInCommon
	return commonM
}

// translate translates a move to the desired language
// Inputs:
// 	-lan:  	Code of the language to translate to
//	-move:	Move that is going to be translated
// Returns:
//	-move: Translated move
func translateMove(lan string, move parser.Node) parser.Node {
	if lan != "en" {
		move = parser.GetMove(move.URL, lan)
	}
	return move
}

// hasMove returns true if a move is a move is on given list
// Inputs:
// 	-m:  	Move to be compared to agaisnt the list of moves
//	-moves:	List of moves in which we are looking for the desired move
// Returns:
//	-true or false depending on if the move was on the list or not
func hasMoves(m string, moves []parser.Move) bool {
	for _, b := range moves {
		if b.Move.Name == m {
			return true
		}
	}
	return false
}

// swap is a function that swaps two given elements of a list.
func swap(i1 int, i2 int, p *[]parser.Pokemon) {
	tmp := (*p)[i1]
	(*p)[i1] = (*p)[i2]
	(*p)[i2] = tmp
}
