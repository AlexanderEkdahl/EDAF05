package main

import (
	"fmt"
	"io"
	"os"
)

type male struct {
	id       int
	name     string
	rankings []int
	next     int
	match    *female
}

type female struct {
	name     string
	rankings map[int]int // lower rank(value) is better
	match    *male
}

type match struct {
	available Stack
	males     map[int]*male
	females   map[int]*female
	n         int
}

func (m *match) match() {
	for m.available.Len() > 0 {
		male := m.available.Pop().(*male)
		wife := m.females[male.rankings[male.next]]
		male.next++

		if wife.match == nil {
			male.match = wife
			wife.match = male
		} else if wife.rankings[male.id] < wife.rankings[wife.match.id] {
			m.available.Push(wife.match)
			male.match = wife
			wife.match = male
		} else {
			m.available.Push(male)
		}
	}
}

func (m *match) print(output io.Writer) {
	for i := 1; i <= m.n*2; i += 2 {
		fmt.Fprintln(output, m.males[i].name, "--", m.males[i].match.name)
	}
}

func main() {
	m := match{
		males:   make(map[int]*male),
		females: make(map[int]*female)}
	m.Parse(os.Stdin)
	m.match()
	m.print(os.Stdout)
}
