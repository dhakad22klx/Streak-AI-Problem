//This is to find the path form source to destination

package main

import (
	"bufio"
	"fmt"
	"os"
	// "reflect"
	// "sync"
	// "sort"
)

const (
	M = 1e9 + 7
)

var in *bufio.Reader
var out *bufio.Writer

func init() {
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
}

func scan(a ...interface{}) {
	fmt.Fscan(in, a...)
}

func scanLine() (string, error) {
	input, err := in.ReadString('\n')
	return input, err
}
func printf(f string, a ...interface{}) {
	// fmt.Fprintf(out, f, a...)
	fmt.Printf(f, a...)
}

//DFS Function

func findPaths(adj map[int][]int, src int, des int, path []int, paths *[][]int) {

	// fmt.Println("Src := ", src)
	if src == des {
		// fmt.Println("path = ", path)
		*paths = append(*paths, path)
		return
	}

	path = append(path, src)

	for _, child := range adj[src] {
		// fmt.Println("Child :=", child)
		findPaths(adj, child, des, path, paths)
	}

}

func main() {

	edges := [][]int{
		{0, 1}, {0, 2}, {1, 2}, {1, 3}, {2, 3}, {3, 4},
	}

	src := 0
	des := 4

	adj := make(map[int][]int)

	var path []int
	var paths [][]int

	for _, edge := range edges {
		u := edge[0]
		v := edge[1]
		adj[u] = append(adj[u], v)
	}

	// fmt.Println("Child of 0", adj[0])

	findPaths(adj, src, des, path, &paths)

	//Check printing paths

	fmt.Println("Total number of paths are : ", len(paths))
	for i, path := range paths {

		fmt.Println("Path", i+1, " := ", path)
	}

}
