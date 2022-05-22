package operation

import (
	"fmt"
	"log"
	"os"
)

func Run(){
	
	a := []byte("dddddd")

	// b := fmt.Printf("%d",a)

	// fmt.Println("b: ",b)

	fmt.Println(a)

	path , err := os.Getwd(); if err != nil {
		log.Println(err)
	}


	file,err := os.OpenFile(path+"/data/test.txt",os.O_WRONLY|os.O_TRUNC|os.O_CREATE,0666); if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close()

	len, err := file.Write(a); if err != nil {
		log.Fatalf("failed writing file: %s", err)
	}
	fmt.Printf("\n file name: %s", file.Name())
	fmt.Printf("\n byte length: %d\n", len)
	
}