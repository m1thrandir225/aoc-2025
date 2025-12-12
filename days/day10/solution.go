package main

import (
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/draffensperger/golp"
	"github.com/m1thrandir225/aoc_2025/util"
)

func part1(input string) any {
	ans := 0

	for _, line := range strings.Split(input, "\n") {
		words := strings.Split(line, " ")

		goal := words[0][1 : len(words[0])-1]

		goalN := 0
		for i, c := range goal {
			if c == '#' {
				goalN += 1 << i
			}
		}

		buttons := words[1 : len(words)-1]

		B := make([]int, 0)
		NS := make([][]int, 0)

		for _, button := range buttons {
			content := button[1 : len(button)-1]

			parts := strings.Split(content, ",")

			ns := make([]int, 0)

			for _, p := range parts {
				v, _ := strconv.Atoi(p)
				ns = append(ns, v)
			}

			buttonN := 0
			for _, x := range ns {
				buttonN += 1 << x
			}

			B = append(B, buttonN)
			NS = append(NS, ns)
		}

		score := len(buttons)

		for a := 0; a < (1 << len(buttons)); a++ {
			an := 0
			aScore := 0

			for i := 0; i < len(buttons); i++ {
				if (a>>i)&1 == 1 {
					an ^= B[i]
					aScore++
				}
			}

			if an == goalN {
				if aScore < score {
					score = aScore
				}
			}
		}

		ans += score
	}

	return ans
}

func part2(input string) any {
	lines := strings.Split(input, "\n")

	total := 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		targetStr := parts[len(parts)-1]
		buttonsStrs := parts[1 : len(parts)-1]

		targets := util.StrToIntArr(targetStr)
		var affects [][]int
		for _, buttonStr := range buttonsStrs {
			affects = append(affects, util.StrToIntArr(buttonStr))
		}

		lp := golp.NewLP(0, len(buttonsStrs))
		obj := make([]float64, len(buttonsStrs))
		for j := range buttonsStrs {
			obj[j] = 1.0
			lp.SetInt(j, true)
		}
		lp.SetObjFn(obj)

		for i := range len(targets) {
			row := make([]float64, len(buttonsStrs))
			hasConstraint := false
			for j := range len(buttonsStrs) {
				if slices.Contains(affects[j], i) {
					row[j] = 1
					hasConstraint = true
				}
			}
			if hasConstraint {
				lp.AddConstraint(row, golp.EQ, float64(targets[i]))
			}
		}

		lp.Solve()

		vars := lp.Variables()
		for j := range buttonsStrs {
			total += int(math.Round(vars[j]))
		}
	}
	return total
}
