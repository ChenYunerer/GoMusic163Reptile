package main

import (
	"os/user"
	"fmt"
	"os"
)

func main() {
	user, err := user.Current()
	if err == nil {
		fmt.Println(user.HomeDir)
	}
	fmt.Println(string(os.PathSeparator))
}
