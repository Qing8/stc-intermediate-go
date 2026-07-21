package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
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

// func main() {
// 	sc := bufio.NewScanner(os.Stdin)
// 	sc.Scan()
// 	kind := sc.Text()
// 	sc.Scan()
// 	dim, _ := strconv.ParseFloat(sc.Text(), 64)
// 	var s interface{ Area() float64 }
// 	switch strings.ToLower(kind) {
// 	case "circle":
// 		s = Circle{dim}
// 	case "square":
// 		s = Square{dim}
// 	}

// 	if s != nil {
// 		fmt.Printf("%.2f\n", s.Area())
// 	}
// }

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	age, err := validateAge(sc.Text())
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	} else {
		fmt.Printf("age: %d\n", age)
	}
}

func validateAge(s string) (int, error) {
	age, err := strconv.Atoi(s)
	if err != nil {
		err = fmt.Errorf("parse: %w", err)
		return age, err
	}
	if age < 0 {
		return age, errors.New("negative")
	}
	return age, nil
}
