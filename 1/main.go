package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("2/input.txt")
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
	counts := map[int]int{}
	for _, val := range lst2 {
		counts[val] += 1
	}
	total := 0
	for _, val := range lst1 {
		times, ok := counts[val]
		if !ok {
			times = 0
		}
		total += val * times
	}
	fmt.Println("Total similarity:", total)
}
