package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	r float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.r * c.r
}

type Square struct {
	r float64
}

func (s Square) Area() float64 {
	return s.r * s.r
}

func main_interface() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	kind := sc.Text()
	sc.Scan()
	dim, _ := strconv.ParseFloat(sc.Text(), 64)
	var s interface{ Area() float64 }
	switch strings.ToLower(kind) {
	case "circle":
		s = Circle{dim}
	case "square":
		s = Square{dim}
	}

	if s != nil {
		fmt.Printf("%.2f\n", s.Area())
	}
}
