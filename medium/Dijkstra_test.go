package medium_test

import (
	"testing"

	"github.com/charlie4284/leetcode/medium"
	"github.com/go-test/deep"
)

var NodeA = medium.Node{
	Label: "A",
	Edges: map[string]*medium.Edge{
		"B": &medium.Edge{
			Distance: 8,
			Node:     &NodeB,
		},
	},
}

var NodeB = medium.Node{
	Label: "B",
	Edges: map[string]*medium.Edge{
		"C": &medium.Edge{
			Distance: 2,
			Node:     &NodeC,
		},
		"D": &medium.Edge{
			Distance: 6,
			Node:     &NodeD,
		},
	},
}

var NodeC = medium.Node{
	Label: "C",
	Edges: map[string]*medium.Edge{
		"D": &medium.Edge{
			Distance: 1,
			Node:     &NodeD,
		},
		"E": &medium.Edge{
			Distance: 6,
			Node:     &NodeE,
		},
	},
}

var NodeD = medium.Node{
	Label: "D",
	Edges: map[string]*medium.Edge{
		"E": &medium.Edge{
			Distance: 2,
			Node:     &NodeE,
		},
	},
}

var NodeE = medium.Node{
	Label: "E",
}

func AssertEqual(t *testing.T, result interface{}, expected interface{}) {
	t.Helper()
	if diff := deep.Equal(result, expected); diff != nil {
		t.Errorf("Unequal results. \nDiff: %v", diff)
	}
}

func Test_ShortestPath(t *testing.T) {
	// testNodes := map[string]*medium.Node{
	// 	"A": &NodeA,
	// 	"B": &NodeB,
	// 	"C": &NodeC,
	// 	"D": &NodeD,
	// 	"E": &NodeE,
	// }
	t.Run("Test get shortest path", func(t *testing.T) {
		result := NodeA.ShortestPath("E")
		expected := []*medium.Node{
			&NodeA, &NodeB, &NodeC, &NodeD, &NodeE,
		}
		AssertEqual(t, result, expected)
	})
}
