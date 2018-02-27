package util

import "fmt"

type Log struct {
}

func (*Log) I(tag string, msg ...interface{}) {
	fmt.Println("[INFO:] ", tag, "-->", msg)
}

func (*Log) E(tag string, msg ...interface{}) {
	fmt.Println("[ERROR:] ", tag, "-->", msg)
}

func (*Log) D(tag string, msg ...interface{}) {
	fmt.Println("[DEBUG:] ", tag, "-->", msg)
}
