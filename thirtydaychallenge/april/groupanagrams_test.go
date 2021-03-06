package april

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	cases := []struct {
		given    []string
		expected [][]string
	}{
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{
				{"eat", "tea", "ate"},
				{"tan", "nat"},
				{"bat"},
			},
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("groupAnagrams(%v)", tc.given), func(t *testing.T) {
			actual := groupAnagrams(tc.given)
			assert.ElementsMatch(t, tc.expected, actual)
		})
	}
}
