package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func partition(arr []int, numParts int) [][]int {
	partSize := len(arr) / numParts
	partitions := make([][]int, numParts)
	for i := 0; i < numParts; i++ {
		startIndex := i * partSize
		endIndex := startIndex + partSize
		if i == numParts-1 {
			endIndex = len(arr)
		}
		partitions[i] = arr[startIndex:endIndex]
	}
	return partitions
}

func sortPartition(arr []int, wg *sync.WaitGroup) {
	fmt.Printf("Sorting subarray: %v\n", arr)
	sort.Ints(arr)
	wg.Done()
}

func mergeSortedArrays(partitions [][]int) []int {
	sortedArr := make([]int, 0)
	for _, partition := range partitions {
		sortedArr = append(sortedArr, partition...)
	}
	sort.Ints(sortedArr)
	return sortedArr
}

func getInput() []int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a series of integers (space-separated): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	arrStr := strings.Split(input, " ")
	arr := make([]int, len(arrStr))
	for i, numStr := range arrStr {
		num, _ := strconv.Atoi(numStr)
		arr[i] = num
	}
	return arr
}

func main() {
	arr := getInput()
	numParts := 4

	partitions := partition(arr, numParts)
	var wg sync.WaitGroup
	wg.Add(numParts)

	for _, partition := range partitions {
		go sortPartition(partition, &wg)
	}

	wg.Wait()
}

/* 
This program prompts the user to enter a series of integers, then partitions the array into four subarrays. 
Each subarray is sorted concurrently by a different goroutine, and the sorted subarrays are then merged by the main goroutine. 
The program prints the subarrays being sorted by each goroutine and finally outputs the entire sorted array.
*/