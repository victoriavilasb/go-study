package main

import (
	"fmt"
	"strings"
)

func main()  {
	var str string
	fmt.Scanf("%s", &str)

	str = strings.ToLower(str)

	if strings.HasPrefix(str, "i") && strings.HasSuffix(str, "n") && strings.ContainsAny(str, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
