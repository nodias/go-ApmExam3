package main

import (
	"flag"
	"fmt"
)

func main() {
	phase := flag.String("phase", "local", "description")
	fmt.Println(*phase)
	flag.Parse()
	fmt.Println(*phase)
}
