package main

import (
	"flag"
	"fmt"
	"os"
)

var bhasad = flag.String("b", "ask help", "sorry, no help")

func main() {
	flag.Parse()
	fmt.Println(*bhasad)
	fmt.Println(os.Getenv(*bhasad))
}
