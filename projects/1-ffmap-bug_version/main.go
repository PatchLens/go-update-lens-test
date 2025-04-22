package main

import (
	"test/pkg"
)

func main() {
	kv, err := pkg.Open("test.ffmap")
	if err != nil {
		panic(err)
	}
	pkg.OperateCSV(kv)
	pkg.OperateInterface(kv)
}
