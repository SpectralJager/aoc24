package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("1/input.txt")
	buff := bytes.NewBuffer(input)
	lst1, lst2 := []int{}, []int{}
	for {
		line, err := buff.ReadString('\n')
		if err != nil {
			break
		}
		line = line[:len(line)-1]
		lineValue := strings.Split(line, "   ")
		val1, _ := strconv.Atoi(lineValue[0])
		val2, _ := strconv.Atoi(lineValue[1])
		lst1 = append(lst1, val1)
		lst2 = append(lst2, val2)
	}
	slices.Sort(lst1)
	slices.Sort(lst2)
	total := 0
	for i := 0; i < len(lst1); i++ {
		dist := (lst1[i] - lst2[i])
		if dist < 0 {
			dist *= -1
		}
		// fmt.Printf("%d - %d = %d\n", lst1[i], lst2[i], dist)
		total += dist
	}
	fmt.Println("Total distance:", total)
}
