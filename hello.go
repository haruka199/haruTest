package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello,世界")
	f, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("error")
	}
	defer f.Close()

	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			break
		}
		fmt.Println(string(buf[:n]))

	}
}
