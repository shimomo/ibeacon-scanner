package main

import (
	"fmt"

	"github.com/shimomo/ibeacon-scanner"
)

func main() {
	ibeacon.Scan(func(uuid string, major string, minor string) {
		fmt.Println(uuid)
		fmt.Println(major)
		fmt.Println(minor)
	}, 1000)
}
