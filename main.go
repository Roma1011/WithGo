package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("C:\\Users\\Student\\Desktop\\tx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	var hello string = "hello World"
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
