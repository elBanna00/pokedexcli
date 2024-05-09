package tests

import (
	"github.com/elBanna00/pokedexcli/repl"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "Hello World",
			expected: []string{
				"hello",
				"world",
			},
		},
	}
	for _, cs := range cases {
		actual := repl.Cleanstr(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("Lenghts DoesNOT match %v , %v", len(cs.expected), len(actual))
			continue
		}
		for i := range actual {
			if actual[i] != cs.expected[i] {
				t.Errorf("%v doesn't equel %v", actual[i], cs.expected[i])
			}
		}
	}
}
