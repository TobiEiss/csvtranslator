package transformer_test

import (
	"log"
	"testing"

	"github.com/TobiEiss/csvtranslator/pkg/transformer"
)

func TestPerrow(t *testing.T) {
	source := [][]string{
		{"header1", "header2"},
		{"entry11", "entry12"},
		{"entry21", "entry22"},
	}

	template := [][]string{
		{"uobj", "val"},
		{"PERSON", "!header1"},
		{"PERSON", "!header2"},
	}

	expected := [][]string{
		{"uobj", "val"},
		{"PERSON", "entry11"},
		{"PERSON", "entry12"},
		{"PERSON", "entry21"},
		{"PERSON", "entry22"},
	}

	result := transformer.TransformPerRow(source, template)
	log.Println(result)
	for i := 0; i < len(expected); i++ {
		for j := 0; j < len(expected[0]); j++ {
			if result[i][j] != expected[i][j] {
				t.Errorf("%s is not equal to %s", result[i][j], expected[i][j])
			}
		}
	}
}
