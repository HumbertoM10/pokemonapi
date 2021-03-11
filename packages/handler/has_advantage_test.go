package handler

import (
	"testing"

	"github.com/HumbertoM10/pokemonapi/packages/parser"
)

// input-result struct type
type TestDataItem struct {
	inputs   []string // inputs to `dmgTaken` function
	result   float32  // result of `dmgTaken` function
	hasError bool     // does `dmgTaken` function returns error
}

// test case for dmgTaken function
func TestDmgTaken(t *testing.T) {

	// input-result data items
	dataItems := []TestDataItem{
		{[]string{"pichu", "gyarados"}, 0.5, false},
		{[]string{"gyarados", "pichu"}, 4, false},
		{[]string{"geodude", "pichu"}, 0, false},
	}

	for _, item := range dataItems {
		pokeArr := parser.GetPokemon(item.inputs)
		dmgRelations := []parser.Drelations{}

		result := dmgTaken(pokeArr[1], dmgRelations) // get result of `dmgTaken` function

		if item.hasError {
			// expected an error
		} else {
			// expected a value
			if result != item.result {
				t.Errorf("dmgTaken() with args %v : FAILED, expected %v but got value '%v'", item.inputs, item.result, result)
			} else {
				t.Logf("dmgTaken() with args %v : PASSED, expected %v and got value '%v'", item.inputs, item.result, result)
			}
		}
	}
}
