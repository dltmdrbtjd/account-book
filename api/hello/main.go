package main

import (
	"fmt"

	"github.com/dltmdrbtjd/account-book/api"
)

func main() {
	message := api.Hello("Lee")
	fmt.Println(message)
}
