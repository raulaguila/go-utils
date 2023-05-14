package main

import (
	"fmt"

	"github.com/raulaguila/go-utils/enum"
)

func main() {
	{ // Enum example
		fmt.Println(enum.Pending)
		fmt.Println(enum.Rejected)
		fmt.Println(enum.Accepted)
		fmt.Println(enum.Test)
	}
}
