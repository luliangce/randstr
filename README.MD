# RandStr

a common way to generate a random string with simple setup

```GO
package main

import (
    "github/luliangce/randstr"
    "log"
)

func main() {
    c := randstr.New(10).DisableLetter().EnableNum().Disablesymbol()
    log.Println(c.String())
}

```

by default:

```Go
const letter = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" //enabled
const number = "0123456789" //disabled
const symbol = `!@#$%^&*()-+=,./;'[]\<>?:"{}|` //disabled
```