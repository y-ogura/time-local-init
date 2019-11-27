package main

import (
	"fmt"
	"time"

	"github.com/y-ogura/time-local-init/locale"
)

func main() {
	fmt.Println("vim-go")
	locale.SetLocale()
	fmt.Println(time.Now())
}
