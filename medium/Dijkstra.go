package medium

import (
	"fmt"
	"sort"
)

// Node - Node of a graph
type Node struct {
	Label string
	Edges map[string]*Edge
}

// Edge - Edge connecting two nodes
type Edge struct {
	Distance int
	Node     *Node
}

// ShortestPath - given label/name of the target node, searches for the shortest path. If none are found, empty slice is returned.
func (n *Node) ShortestPath(target string) []*Node {
	distanceMap := map[string]Edge{
		n.Label: Edge{
			Distance: 0,
			Node:     n,
		},
	}
	visitedMap := map[string]bool{}
	toVisitNodes := []*Edge{}

	//Initialize toVisitNodes
	for _, edge := range n.Edges {
		toVisitNodes = append(toVisitNodes, edge)
		distanceMap[edge.Node.Label] = Edge{
			Distance: edge.Distance,
			Node:     n,
		}
	}
	sort.Slice(toVisitNodes, func(i, j int) bool {
		return toVisitNodes[i].Distance < toVisitNodes[j].Distance
	})

FindTarget:
	for len(toVisitNodes) != 0 {
		curNode := getClosestUnvisitedNode(toVisitNodes, visitedMap)
		fmt.Println(curNode.Node.Label)

		for nodeLabel, edge := range curNode.Node.Edges {
			if _, ok := visitedMap[nodeLabel]; ok {
				continue
			}
			toVisitNodes = append(toVisitNodes, curNode.Node.Edges[nodeLabel])
			sort.Slice(toVisitNodes, func(i, j int) bool {
				return toVisitNodes[i].Distance < toVisitNodes[j].Distance
			})
			if distEdge, ok := distanceMap[nodeLabel]; !ok {
				distanceMap[nodeLabel] = Edge{
					Distance: edge.Distance,
					Node:     curNode.Node,
				}
			} else if distEdge.Distance > edge.Distance+distanceMap[curNode.Node.Label].Distance {
				distanceMap[nodeLabel] = Edge{
					Distance: edge.Distance + distanceMap[curNode.Node.Label].Distance,
					Node:     curNode.Node,
				}
			}
		}
		if curNode.Node.Label == target {
			// TODO: stuff
			break FindTarget
		}
	}

	parentNode, ok := distanceMap[target]
	if !ok {
		return nil
	}
	reversedSlice := []*Node{parentNode.Node.Edges[target].Node}
	for parentNode.Node.Label != n.Label {
		reversedSlice = append(reversedSlice, parentNode.Node)
		parentNode = distanceMap[parentNode.Node.Label]
	}
	reversedSlice = append(reversedSlice, parentNode.Node)
	for i, j := 0, len(reversedSlice)-1; i < j; i, j = i+1, j-1 {
		reversedSlice[i], reversedSlice[j] = reversedSlice[j], reversedSlice[i]
	}
	return reversedSlice
}

func getClosestUnvisitedNode(sortedEdges []*Edge, visitedMap map[string]bool) *Edge {
	for edgeIdx, edge := range sortedEdges {
		if _, ok := visitedMap[edge.Node.Label]; !ok {
			visitedMap[edge.Node.Label] = true
			return sortedEdges[edgeIdx]
		}
	}
	return nil
}
