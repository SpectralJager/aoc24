package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ToInts(strs []string) []int {
	ints := []int{}
	for _, str := range strs {
		v, _ := strconv.Atoi(str)
		ints = append(ints, v)
	}
	return ints
}

func cut[T any](arr []T, index int) []T {
	newArr := make([]T, len(arr))
	copy(newArr, arr)
	return append(newArr[:index], newArr[index+1:]...)
}

func every[T any](arr []T, callback func(i int, arr []T) bool) bool {
	newArr := make([]T, len(arr))
	copy(newArr, arr)
	for i := range newArr {
		if !callback(i, arr) {
			return false
		}
	}
	return true
}

func some[T any](arr []T, callback func(i int, arr []T) bool) bool {
	newArr := make([]T, len(arr))
	copy(newArr, arr)
	for i := range newArr {
		if callback(i, arr) {
			return true
		}
	}
	return false
}

func everyOnce[T any](arr []T, callback func(i int, arr []T) bool) bool {
	newArr := make([]T, len(arr))
	copy(newArr, arr)
	if every(newArr, callback) {
		return true
	}
	return some(newArr, func(i int, arr []T) bool {
		return every(cut(arr, i), callback)
	})
}

func main() {
	data, _ := os.ReadFile("2/input.txt")
	rd := bufio.NewReader(bytes.NewBuffer(data))
	total1 := 0
	total2 := 0
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			break
		}
		line = line[:len(line)-1]
		nums := ToInts(strings.Split(line, " "))
		if every(nums, func(i int, arr []int) bool {
			return i == 0 || (arr[i] < arr[i-1] && arr[i-1]-arr[i] < 4)
		}) || every(nums, func(i int, arr []int) bool {
			return i == 0 || (arr[i] > arr[i-1] && arr[i]-arr[i-1] < 4)
		}) {
			total1 += 1
		}
		if everyOnce(nums, func(i int, arr []int) bool {
			return i == 0 || (arr[i] < arr[i-1] && arr[i-1]-arr[i] < 4)
		}) || everyOnce(nums, func(i int, arr []int) bool {
			return i == 0 || (arr[i] > arr[i-1] && arr[i]-arr[i-1] < 4)
		}) {
			total2 += 1
		}
	}
	fmt.Println("Total 1: ", total1)
	fmt.Println("Total 2: ", total2)
}
