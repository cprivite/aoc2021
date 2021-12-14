package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/echojc/aocutil"
)

func main() {
	//Read Input
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	input, err := i.Strings(2021, 12)
	if err != nil {
		log.Fatal(err)
	}

	directions := [][2]string{}
	for _, line := range input {
		letters := [2]string{strings.Split(string(line), "-")[0], strings.Split(string(line), "-")[1]}
		directions = append(directions, letters)

	}

	adj := make(map[string][]string)
	paths := [][]string{}
	path := make([]string, 0)
	visited := make(map[string]bool)

	for _, edge := range directions {
		v := edge[0]
		w := edge[1]
		if v != "end" && w != "start" {
			addEdge(v, w, adj)
		}
		if w != "end" && v != "start" {
			addEdge(w, v, adj)
		}
	}

	fmt.Println(adj)
	DFS("start", adj, &paths, path, visited, 1)

	for _, i := range paths {
		fmt.Println(i)
	}

	fmt.Println(len(paths))

}

func addEdge(v string, w string, adj map[string][]string) {
	edge := adj[v]
	edge = append(edge, w)
	adj[v] = edge

}

func DFS(v string, adj map[string][]string, paths *[][]string, path []string, visited map[string]bool, twice int) {
	newPath := make([]string, len(path))
	copy(newPath, path)

	newPath = append(newPath, v)
	if v != strings.ToUpper(v) && v != "start" {
		visited[v] = true
	}

	if v != strings.ToUpper(v) && strTains(newPath, v) > 1 && v != "start" {
		twice = 0
	}

	if v == "end" {
		newPath = append(path, "end")
		*paths = append((*paths), newPath)
	} else {
		for i := 0; i < len(adj[v]); i++ {
			if adj[v][i] == strings.ToUpper(adj[v][i]) || adj[v][i] == "end" {
				DFS(adj[v][i], adj, paths, newPath, visited, twice)
			} else if strTains(newPath, adj[v][i]) < 1 || twice > 0 {
				DFS(adj[v][i], adj, paths, newPath, visited, twice)
			}

		}
	}

	visited[v] = false
	twice = 1
}

func strTains(myString []string, yourString string) int {
	contains := 0
	for i := range myString {
		if strings.Contains(myString[i], yourString) {
			contains++
		}
	}
	return contains
}
