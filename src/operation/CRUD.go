package operation

import (
	"fmt"
	"log"
	"os"
)

func Run(){
	fmt.Println("run")
	
	// a := []byte("hello world")

	file,err := os.Create("test.txt"); if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close()

	len, err := file.WriteString("helloo"); if err != nil {
		log.Fatalf("failed writing file: %s", err)
	}
	fmt.Printf("\n file name: %s", file.Name())
	fmt.Printf("\n byte length: %d\n", len)
	
}