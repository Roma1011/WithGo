package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.OpenFile("C:\\Users\\Student\\Desktop\\tx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		_ = fmt.Errorf("Error opening file:", err)
		return
	}
	var result int = 0
	result, err = divide(5, 3)
	if err != nil {
		_ = fmt.Errorf("Error opening file:", err)
		return
	}
	var hello string = strconv.Itoa(result)
	arry := make([]byte, len(hello))

	for i := 0; i < len(hello); i++ {
		arry[i] = byte(hello[i])
	}

	_, err = file.Write(arry)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	_ = file.Close()
}

func divide(i, j int) (int int, err error) {
	if i == 0 || j == 0 {
		err = errors.New("cannot divide by zero")
	}
	return i / j, err
}
