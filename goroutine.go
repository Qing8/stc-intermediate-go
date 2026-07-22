package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main_goroutine() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	fields := strings.Fields(sc.Text())
	nums := make([]int, n)
	for i, f := range fields {
		val, err := strconv.Atoi(f)
		nums[i] = val
		if err != nil {
			panic(err.Error())
		}
	}

	length := len(nums)
	numGroups := (length + 3) / 4 // Get the ceiling of integer division
	var total int
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < numGroups; i++ {
		var curBatch []int
		start, end := i*4, (i+1)*4
		if end > length {
			end = length
		}
		for j := start; j < end; j++ {
			curBatch = append(curBatch, nums[j])
		}
		wg.Add(1)
		go func(total *int, curBatch []int) {
			defer wg.Done()
			curBatchTotal := 0
			for l := 0; l < len(curBatch); l++ {

				curBatchTotal += curBatch[l]
			}
			mu.Lock()
			*total += curBatchTotal
			mu.Unlock()
		}(&total, curBatch)
	}
	wg.Wait()

	fmt.Println(total) // replace with the real total
}
