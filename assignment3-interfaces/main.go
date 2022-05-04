package main

import (
	"fmt"
	"os"
)


func main() {
	f, err := os.Create("myfile.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	w, err := f.WriteString("Hey this is Myfile!!!")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(w, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}