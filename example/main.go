package main

import (
	"log"

	"github.com/luliangce/randstr"
)

func main() {
	c := randstr.New(10).DisableLetter().EnableNum().Disablesymbol()
	log.Println(c.String())
}
