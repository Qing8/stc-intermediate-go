package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main_error() {
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
