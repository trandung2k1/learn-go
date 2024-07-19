package main

import (
	"fmt"
	"packages-modules/configs"
)

func main() {
	configs.Config()
	fmt.Println(configs.Url)
}
