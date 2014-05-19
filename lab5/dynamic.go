package main

import (
	"fmt"
	"os"
)

const (
	gap1 = iota
	gap2
	copy
)

var scores = make(map[string]int)

type subject struct {
	name  string
	value string
}

func matchScore(a, b uint8) int {
	return scores[string([]uint8{a, b})]
}

func reverse(s []uint8) string {
	o := make([]uint8, len(s))
	i := len(o)
	for _, c := range s {
		i--
		o[i] = c
	}
	return string(o)
}

func align(s1, s2 string) (string, string, int) {
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

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			max := d[i-1][j] + matchScore(s2[j-1], '*')

			if s := d[i][j-1] + matchScore('*', s1[i-1]); max < s {
				max = s
			}

			if s := d[i-1][j-1] + matchScore(s1[i-1], s2[j-1]); max < s {
				max = s
			}

			d[i][j] = max
		}
	}

	// for i := range d {
	// 	fmt.Println(d[i])
	// }

	a := make([]uint8, 0)
	b := make([]uint8, 0)
	i, j := len(s1), len(s2)

	for i > 0 || j > 0 {
		if i > 0 && j > 0 && (d[i][j] == (d[i-1][j-1] + matchScore(s1[i-1], s2[j-1]))) {
			a = append(a, s1[i-1])
			b = append(b, s2[j-1])
			i--
			j--
		} else if i > 0 && (d[i][j] == d[i-1][j]+matchScore(s1[i-1], '*')) {
			a = append(a, s1[i-1])
			b = append(b, '-')
			i--
		} else if j > 0 && (d[i][j] == d[i][j-1]+matchScore('*', s2[j-1])) {
			a = append(a, '-')
			b = append(b, s2[j-1])
			j--
		}
	}

	return reverse(a), reverse(b), d[len(s1)][len(s2)]
}

func main() {
	file, _ := os.Open(os.Args[1])
	scores = parseScores(file)

	file, _ = os.Open(os.Args[2])
	subjects := parseSubjects(file)

	for i := 0; i < len(subjects); i++ {
		for j := i + 1; j < len(subjects); j++ {
			a, b, s := align(subjects[i].value, subjects[j].value)
			fmt.Printf("%v--%v: %v\n%v\n%v\n", subjects[i].name, subjects[j].name, s, a, b)
		}
	}
}
