package main

import "fmt"

type node struct {
	provName  string
	neighbors []string
	visited   bool
	inDegree  int
}

func main() {
	var depList = map[string][]string{
		"ui":        []string{},
		"auth":      []string{},
		"product":   []string{},
		"mongodb":   []string{},
		"order":     []string{"product", "mongodb"},
		"api-gw":    []string{"auth", "product", "order"},
		"order-sim": []string{"order"},
	}

	var queue []string
	for proc, deps := range depList {
		if len(deps) == 0 {
			queue = append(queue, proc)
			delete(depList, proc)
		}
	}

	for len(queue) != 0 {
		dependency := queue[0]
		queue = queue[1:]
		fmt.Printf("start process: %s\n", dependency)

		for proc, deps := range depList {
			found, newDeps := foundAndRemoveDep(deps, dependency)
			if found {
				if len(newDeps) == 0 {
					queue = append(queue, proc)
					delete(depList, proc)
				} else {
					depList[proc] = newDeps
				}
			}
		}
	}

	if len(depList) != 0 {
		fmt.Println("deadlock...")
		fmt.Println(depList)
	}
}

func foundAndRemoveDep(depList []string, dep string) (bool, []string) {
	var newList []string
	var found bool
	for _, val := range depList {
		if val == dep {
			found = true
			continue
		}

		newList = append(newList, val)
	}

	return found, newList
}
