package main

import (
	"fmt"
	"github.com/VladMak/golearnprojects/internal/gpl"
	"os"
)

func main() {
	fmt.Println("Project Go Programming language")
	gpl.Dup(os.Args[1:])
}