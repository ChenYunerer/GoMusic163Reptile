package main

import (
	"os/user"
	"fmt"
	"os"
)

func main() {
	fmt.Println("123")
	fmt.Println("1233333")
}

func getUserHome() {
	user, err := user.Current()
	if err == nil {
		fmt.Println(user.HomeDir)
	}
}

func getPathSeparator() {
	fmt.Println(string(os.PathSeparator))
}


