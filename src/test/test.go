package main

import (
	"os/user"
	"fmt"
	"os"
)

func main() {
	ids := make([]int, 0)
	fmt.Println(ids)
	for i := 0; i < 21; i++ {
		ids = append(ids, i)
	}
	fmt.Println(ids)
	for i := 0; i < 5; i++ {
		fmt.Println(5 / 4)
		id1 := ids[len(ids)*i: len(ids)/5*(i+1)]
		fmt.Println(id1)
	}
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
