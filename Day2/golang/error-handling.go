package main

import (
	"fmt"
	"os"
	"os/user"
	"github.com/sttk/stringcase"
)

func main() {
	u, err := user.Current()

	if err != nil {
		fmt.Println("Cannot get current user:", err)
		os.Exit(1)
	}

	fmt.Printf("Hello %s, welcome !\n", stringcase.PascalCase(u.Username) )
}
