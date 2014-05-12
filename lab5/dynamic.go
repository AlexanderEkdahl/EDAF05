package main

import (
	"fmt"
	"os"
)

var scores = make(map[string]int)
var subjects = make([]*subject, 3)

type subject struct {
	name  string
	value string
}

func matchScore(a, b uint8) int {
	return scores[string([]uint8{a, b})]
}

func gapScore(r uint8) int {
	return matchScore(r, uint8('*'))
}

func align(s1, s2 string) (string, int) {
	d := make([][]int, len(s1))
	for i := range d {
		d[i] = make([]int, len(s2))
	}

	for i := 1; i < len(s1); i++ {
		for j := 1; j < len(s2); j++ {
			max := d[i-1][j] + gapScore(s2[j-1])

			if s := d[i][j-1] + gapScore(s1[i-1]); max < s {
				max = s
			}

			if s := d[i-1][j-1] + matchScore(s1[i-1], s2[j-1]); max < s {
				max = s
			}

			d[i][j] = max
		}
	}

	for i := range d {
		fmt.Println(d[i])
	}

	return "*", d[len(s1)-1][len(s2)-1]
}

func main() {
	subjects[0] = &subject{"Sphinx", "KQRK"}
	subjects[1] = &subject{"Bandersnatch", "KAK"}
	subjects[2] = &subject{"Snark", "KQRIKAAKABK"}

	// scores["KK"] = 5
	// scores["KQ"] = 1
	// scores["KR"] = 2
	// scores["KA"] = -1
	// scores["KI"] = -3
	// scores["KB"] = 0
	// scores["K*"] = -4
	// scores["QQ"] = 5
	// scores["QR"] = 1
	// scores["QA"] = -1
	// scores["QI"] = -3
	// scores["QB"] = 0
	// scores["Q*"] = -4
	// scores["RR"] = 5
	// scores["RA"] = -1
	// scores["RI"] = -3
	// scores["RB"] = -1
	// scores["R*"] = -4
	// scores["AA"] = 4
	// scores["AI"] = -1
	// scores["AB"] = -2
	// scores["A*"] = -4
	// scores["II"] = 4
	// scores["IB"] = -3
	// scores["I*"] = -4
	// scores["BB"] = 4
	// scores["B*"] = -4
	// scores["**"] = 1

	file, _ := os.Open("fixtures/BLOSUM62.txt")
	scores = parseScores(file)

	for i := 0; i < len(subjects); i++ {
		for j := i + 1; j < len(subjects); j++ {
			r, s := align(subjects[i].value, subjects[j].value)
			fmt.Printf("%v--%v: %v\n%v\n", subjects[i].name, subjects[j].name, s, r)
		}
	}
}
