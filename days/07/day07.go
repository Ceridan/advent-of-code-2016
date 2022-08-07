package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Part1(ips []string) int {
	rb := regexp.MustCompile("(\\[[a-z]+\\])")
	rn := regexp.MustCompile("([a-z]+)")
	tls := 0

	for _, ip := range ips {
		brackets := rb.FindAllString(ip, -1)
		bracketsAbba := false
		for _, bracket := range brackets {
			if searchAbba(bracket) {
				bracketsAbba = true
				break
			}
		}

		if bracketsAbba {
			continue
		}

		nonBracketsIp := rb.ReplaceAllString(ip, "|")
		nonBrackets := rn.FindAllString(nonBracketsIp, -1)
		for _, nonBracket := range nonBrackets {
			if searchAbba(nonBracket) {
				tls += 1
				break
			}
		}
	}
	return tls
}

func Part2(ips []string) int {
	rb := regexp.MustCompile("(\\[[a-z]+\\])")
	rn := regexp.MustCompile("([a-z]+)")
	ssl := 0

	for _, ip := range ips {
		brackets := rb.FindAllString(ip, -1)
		abaFromBab := make(map[string]struct{})
		for _, bracket := range brackets {
			babs := listAba(bracket)
			for _, bab := range babs {
				aba := convertBabToAba(bab)
				abaFromBab[aba] = struct{}{}
			}
		}

		if len(abaFromBab) == 0 {
			continue
		}

		nonBracketsIp := rb.ReplaceAllString(ip, "|")
		nonBrackets := rn.FindAllString(nonBracketsIp, -1)
	NonBracket:
		for _, nonBracket := range nonBrackets {
			abas := listAba(nonBracket)
			for _, aba := range abas {
				if _, ok := abaFromBab[aba]; ok {
					ssl += 1
					break NonBracket
				}
			}
		}
	}
	return ssl
}

func searchAbba(ip string) bool {
	for i := 0; i <= len(ip)-4; i++ {
		a1, b1, b2, a2 := ip[i], ip[i+1], ip[i+2], ip[i+3]
		if a1 == a2 && b1 == b2 && a1 != b1 {
			return true
		}
	}

	return false
}

func listAba(ip string) []string {
	ipRunes := []rune(ip)
	var abas []string
	for i := 0; i <= len(ip)-3; i++ {
		a1, b, a2 := ipRunes[i], ipRunes[i+1], ipRunes[i+2]
		if a1 == a2 && a1 != b {
			abas = append(abas, string([]rune{a1, b, a2}))
		}
	}

	return abas
}

func convertBabToAba(bab string) string {
	babRunes := []rune(bab)
	return string([]rune{babRunes[1], babRunes[0], babRunes[1]})
}

func main() {
	input, err := os.ReadFile("days/07/input.txt")
	if err != nil {
		panic(err)
	}

	var ips []string
	for _, in := range strings.Split(string(input), "\n") {
		if in == "" {
			continue
		}
		ips = append(ips, strings.Trim(in, " "))
	}

	fmt.Printf("Day 07, part 1: %v\n", Part1(ips))
	fmt.Printf("Day 07, part 2: %v\n", Part2(ips))
}
