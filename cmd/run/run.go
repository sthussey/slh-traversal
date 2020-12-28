package main

import (
	"flag"
	"fmt"
)

func main() {
	var n int
	flag.IntVar(&n, "n", 2, "Dimension")
	flag.Parse()
	graph := computeTopo(n)
	countCache := make(map[GraphCoord]uint64)
	paths := countPaths(graph, countCache)
	fmt.Printf("%d x %d topography has %d unique traversal paths from 1,1.\n", n, n, paths + 1)
}

type GraphCoord struct {
	c int
	r int
}

type GraphNode struct {
	GraphCoord
	down *GraphNode
	right *GraphNode
}

func computeTopo(dim int) *GraphNode {
	root := &GraphNode{GraphCoord: GraphCoord{c: 1, r: 1}, down: nil, right: nil}
	var up *GraphNode
	var left *GraphNode
	var row *GraphNode = root
	for r := 1; r <= dim; r++{
		for c := 1; c <= dim; c++{
			if(c == 1 && r == 1){
				left = row
				continue
			}
			newNode := &GraphNode{GraphCoord: GraphCoord{r: r, c: c}, down: nil, right: nil}
			if(c == 1){
				row = newNode
			}
			if(up != nil){
				up.down = newNode
				up = up.right
			}
			if(left != nil){
				left.right = newNode
			}
			left = newNode
		}
		up = row
		left = nil
	}
	return root
}

func countPaths(root *GraphNode, cache map[GraphCoord]uint64) uint64 {
	if (root == nil){
		return 0
	}
	cacheResult, ok := cache[root.GraphCoord]
	if(ok){
		return cacheResult
	}
	var addlPath uint64 = 0
	if (root.right != nil && root.down != nil){
		addlPath = 1
	}
	paths := countPaths(root.right, cache)
        paths = paths + countPaths(root.down, cache)
	cache[root.GraphCoord] = paths + addlPath
	return paths + addlPath
}

