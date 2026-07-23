package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		nums[i], _ = strconv.Atoi(sc.Text())
	}

	var wg sync.WaitGroup
	total := 0

	length := len(nums)
	nGroup := (length + 3) / 4
	partialSum := make(chan int, nGroup)
	for i := 0; i < nGroup; i++ {

		start, end := i*4, (i+1)*4
		if end > len(nums) {
			end = len(nums)
		}
		wg.Add(1)
		go sum(nums[start:end], partialSum, &wg)
	}
	wg.Wait()
	close(partialSum)

	for v := range partialSum {
		total += v
	}
	fmt.Println(total)
}

func sum(a []int, partialSum chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	partialSum <- s
}
