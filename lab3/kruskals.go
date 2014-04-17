package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type disjoint struct {
	parent *disjoint
	value  string
	rank   int
}

func makeSet(value string) *disjoint {
	e := new(disjoint)
	e.parent = e
	e.value = value
	e.rank = 0
	return e
}

func find(x *disjoint) *disjoint {
	if x.parent != x {
		x.parent = find(x.parent)
	}
	return x.parent
}

func union(x, y *disjoint) {
	root1 := find(x)
	root2 := find(y)

	if root1 == root2 {
		return
	}

	if root1.rank < root2.rank {
		root1.parent = root2
	} else if root1.rank > root2.rank {
		root2.parent = root1
	} else {
		root2.parent = root1
		root1.rank = root1.rank + 1
	}
}

type edge struct {
	a, b   *disjoint
	weight int
}

type edgesSorter []edge

func (e edgesSorter) Len() int           { return len(e) }
func (e edgesSorter) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e edgesSorter) Less(i, j int) bool { return e[i].weight < e[j].weight }

func parse(r io.Reader) (e edgesSorter) {
	var distance = regexp.MustCompile(`(\"[\w, -]+\"|\w+)(?:--(\"[\w, -]+\"|\w+)\s\[(\d+)\])?`)
	scanner := bufio.NewScanner(r)
	cities := make(map[string]*disjoint)

	for scanner.Scan() {
		if match := distance.FindStringSubmatch(scanner.Text()); match[2] != "" {
			distance, _ := strconv.Atoi(match[3])
			e = append(e, edge{cities[match[1]], cities[match[2]], distance})
		} else {
			cities[match[1]] = makeSet(match[1])
		}
	}

	return
}

func main() {
	result := 0
	e := parse(os.Stdin)

	sort.Sort(e)

	for _, edge := range e {
		if find(edge.a) != find(edge.b) {
			result += edge.weight
			union(edge.a, edge.b)
		}
	}

	fmt.Println(result)
}
