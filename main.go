package main

import (
	"fmt"
	"gech/hello"
	"gech/hello/hi"

	"log"
	"os"
)

func main() {
	fmt.Println("main")
	hello.Hello()
	hi.Hi()
	a := []byte("hello world")
	fmt.Println(string(a[0]))

	
	file, err := os.Create("test.txt"); if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	defer file.Close()

	len, err := file.WriteString("hello world"); if err != nil {
		log.Fatalf("failed writing file: %s", err)
	}

	fmt.Printf("\n file name: %s", file.Name())
	fmt.Printf("\n byte length: %d\n", len)

}