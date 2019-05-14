package main

import (
	"github/luliangce/randstr"
	"log"
)

func main() {
	c := randstr.New(10).DisableLetter().EnableNum().Disablesymbol()
	log.Println(c.String())
}
