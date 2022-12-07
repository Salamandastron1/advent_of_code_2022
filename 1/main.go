package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type elf struct {
	ID       int
	calories int
}

func main() {
	elves := []elf{}
	count := 1
	max := 0

	f, err := os.Open("elf_food.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var e elf
	e.ID = 1
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			count++
			elves = append(elves, e)
			e = elf{ID: count}
			// skip blank lines
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		e.calories += num
	}
	elves = append(elves, e)
	sort.Slice(elves, func(i, j int) bool {
		return elves[i].calories < elves[j].calories
	})
	maxThree := elves[len(elves)-3:]
	for _, v := range maxThree {
		max += v.calories
	}
	fmt.Printf("Total calories: %v\nElves: %v", max, maxThree)
}

// 50 stars by 12/25
// calories per elf
// each line with a number is considered the total calories for a single food item
// elves group thier foods by separting each elf with a blank new line
