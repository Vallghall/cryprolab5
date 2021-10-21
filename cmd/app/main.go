package main

import (
	"cryprolab5/internal/hashes"
	"fmt"
	"os"
)

func main() {
	fmt.Println(hashes.HashWithAllFunctions(os.Args[1]))
}
