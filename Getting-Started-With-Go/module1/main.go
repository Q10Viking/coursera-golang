package main

import (
	"fmt"
	"time"
)

const Layout string = "2006-01-02"

func main() {
	msg := "Hello World"
	time := time.Now().Format(Layout)
	fmt.Println(msg, "At", time)
}

/**output
Hello World At 2023-03-23
*/
