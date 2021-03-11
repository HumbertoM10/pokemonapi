package handler

import (
	"testing"

	"github.com/HumbertoM10/pokemonapi/packages/parser"
)

// input-result struct type
type TestDataItem struct {
	pokemon      parser.Pokemon // Inputs to `dmgTaken` function
	dmgRelations []parser.DmgRelations
	result       float32 // Result of `dmgTaken` function
	hasError     bool    // Does `dmgTaken` function returns error
}

// test case for dmgTaken function
func TestDmgTaken(t *testing.T) {

	dataItems := []TestDataItem{
		{
			pokemon: parser.Pokemon{
				Name: "gyarados",
				Types: []parser.PokeType{
					{
						parser.Node{Name: "water"},
					},
					{
						parser.Node{Name: "flying"},
					},
				},
			},
			dmgRelations: []parser.DmgRelations{
				{
					Name: "electric",
					TypeRelation: parser.TypeRelation{
						HalfDmgFrom: []parser.Node{
							{
								Name: "flying",
							},
						},
					},
				},
			},
			result:   0.5,
			hasError: false,
		},
	}

	for _, item := range dataItems {

		result := dmgTaken(item.pokemon, item.dmgRelations) // get result of `dmgTaken` function

		if item.hasError {
			// expected an error
		} else {
			// expected a value
			if result != item.result {
				t.Errorf("dmgTaken(): FAILED, expected %v but got value '%v'", item.result, result)
			} else {
				t.Logf("dmgTaken(): PASSED, expected %v and got value '%v'", item.result, result)
			}
		}
	}
}

// test case for dmgTaken function
func TestDmgDone(t *testing.T) {

	// input-result data items
	dataItems := []TestDataItem{
		{
			pokemon: parser.Pokemon{
				Name: "gyarados",
				Types: []parser.PokeType{
					{
						parser.Node{Name: "water"},
					},
					{
						parser.Node{Name: "flying"},
					},
				},
			},
			dmgRelations: []parser.DmgRelations{
				{
					Name: "electric",
					TypeRelation: parser.TypeRelation{
						DoubleDmgTo: []parser.Node{
							{
								Name: "water",
							},
							{
								Name: "flying",
							},
						},
					},
				},
			},
			result:   4,
			hasError: false,
		},
	}

	for _, item := range dataItems {

		result := dmgDone(item.dmgRelations, item.pokemon) // get result of `dmgDone` function

		if item.hasError {
			// expected an error
		} else {
			// expected a value
			if result != item.result {
				t.Errorf("dmgTaken(): FAILED, expected %v but got value '%v'", item.result, result)
			} else {
				t.Logf("dmgTaken(): PASSED, expected %v and got value '%v'", item.result, result)
			}
		}
	}
}
