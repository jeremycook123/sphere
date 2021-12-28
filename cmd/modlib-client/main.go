package main

import (
	"fmt"

	"github.com/jeremycook123/sphere"
	"github.com/jeremycook123/sphere/pkg/clientlib"
)

func main() {
	fmt.Println("Running client")
	fmt.Println("Config:", sphere.Config())
	fmt.Println(clientlib.Hello())
}
