package main

import (
	"example.com/m/serverProvider"
	"fmt"
)

func main() {

	fmt.Println("running1")

	srv := serverProvider.SrvInit()

	fmt.Println("running2")

	srv.LocalCacheMain()

	fmt.Println("running3")

}
