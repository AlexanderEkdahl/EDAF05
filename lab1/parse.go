package main

import (
	"io"
	"strconv"
	"text/scanner"
)

type parser struct {
	m *match
	s scanner.Scanner
}

type stateFn func(*parser) stateFn

func parseComment(p *parser) stateFn {
	if p.s.Peek() == 'n' {
		return parseN
	}

	for p.s.Scan() != '\n' {
	}
	return parseComment
}

func parseN(p *parser) stateFn {
	p.s.Scan()
	p.s.Scan()
	p.s.Scan()
	n, _ := strconv.Atoi(p.s.TokenText())
	p.m.n = n
	p.s.Scan()
	return parseMale
}

func parseMale(p *parser) stateFn {
	p.s.Scan()
	n, _ := strconv.Atoi(p.s.TokenText())
	name := ""
	for p.s.Scan() != '\n' {
		name = name + p.s.TokenText()
	}
	male := &male{id: n, name: name, rankings: make([]int, p.m.n)}
	p.m.males[n] = male
	p.m.available.Push(male)

	return parseFemale
}

func parseFemale(p *parser) stateFn {
	p.s.Scan()
	n, _ := strconv.Atoi(p.s.TokenText())

	name := ""
	for p.s.Scan() != '\n' {
		name = name + p.s.TokenText()
	}
	p.m.females[n] = &female{name: name, rankings: make(map[int]int)}

	if p.s.Peek() == '\n' {
		p.s.Scan()
		return parseMaleRanking
	}

	return parseMale
}

func parseMaleRanking(p *parser) stateFn {
	p.s.Scan()
	n, _ := strconv.Atoi(p.s.TokenText())
	p.s.Scan()

	i := 0
	for p.s.Scan() != '\n' {
		rank, _ := strconv.Atoi(p.s.TokenText())
		p.m.males[n].rankings[i] = rank
		i++
	}

	return parseFemaleRanking
}

func parseFemaleRanking(p *parser) stateFn {
	p.s.Scan()
	n, _ := strconv.Atoi(p.s.TokenText())
	p.s.Scan()

	rank := 0
	for p.s.Scan() != '\n' {
		i, _ := strconv.Atoi(p.s.TokenText())
		p.m.females[n].rankings[i] = rank
		rank++
	}

	if p.s.Peek() == scanner.EOF {
		return nil
	}

	return parseMaleRanking
}

func (m *match) Parse(input io.Reader) {
	p := &parser{m: m}
	p.s.Init(input)
	p.s.Mode = scanner.ScanInts
	p.s.Whitespace = 1 << ' '

	for state := parseComment; state != nil; {
		state = state(p)
	}
}
