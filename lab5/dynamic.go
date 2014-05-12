package main

import (
	"fmt"
	"os"
)

var scores = make(map[string]int)

type subject struct {
	name  string
	value string
}

func matchScore(a, b uint8) int {
	return scores[string([]uint8{a, b})]
}

func align(s1, s2 string) (string, int) {
	d := make([][]int, len(s1)+1)
	for i := range d {
		d[i] = make([]int, len(s2)+1)
	}

	for i := 1; i < len(s1)+1; i++ {
		d[i][0] = d[i-1][0] + matchScore('*', s1[i-1])
	}

	for i := 1; i < len(s2)+1; i++ {
		d[0][i] = d[0][i-1] + matchScore(s2[i-1], '*')
	}

	a := make([][]string, len(s1)+1)
	for i := range a {
		a[i] = make([]string, len(s2)+1)
	}

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			max := d[i-1][j] + matchScore(s2[j-1], '*')
			c := "N"

			if s := d[i][j-1] + matchScore('*', s1[i-1]); max < s {
				max = s
				c = "W"
			}

			if s := d[i-1][j-1] + matchScore(s1[i-1], s2[j-1]); max < s {
				max = s
				c = "NW"
			}

			d[i][j] = max
			a[i][j] = c
		}
	}

	// for i := range d {
	// 	fmt.Println(d[i])
	// }
	//
	// for i := range d {
	// 	fmt.Println(a[i])
	// }

	return "", d[len(s1)][len(s2)]
}

func main() {
	file, _ := os.Open(os.Args[1])
	scores = parseScores(file)

	file, _ = os.Open(os.Args[2])
	subjects := parseSubjects(file)

	for i := 0; i < len(subjects); i++ {
		for j := i + 1; j < len(subjects); j++ {
			r, s := align(subjects[i].value, subjects[j].value)
			fmt.Printf("%v--%v: %v\n%v", subjects[i].name, subjects[j].name, s, r)
		}
	}
}
