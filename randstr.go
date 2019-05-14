package randstr

import (
	"math/rand"
	"time"
)

const letter = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const number = "0123456789"
const symbol = `!@#$%^&*()-+=,./;'[]\<>?:"{}|`

//SetSeed warp the math/rand.Seed()
//you can also call it directly
func SetSeed(seed int64) {
	rand.Seed(seed)
}

//Config define the way to gen a randstr
type Config struct {
	length int
	letter bool
	number bool
	symbol bool
}

//EnableNum turn on nums
func (c *Config) EnableNum() *Config {
	c.number = true
	return c
}

//DisableNum turn off nums
func (c *Config) DisableNum() *Config {
	c.letter = false
	return c
}

//EnableLetter turn on letters
func (c *Config) EnableLetter() *Config {
	c.letter = true
	return c
}

//DisableLetter turn off letters
func (c *Config) DisableLetter() *Config {
	c.letter = false
	return c
}

//Enablesymbol turn on symbols
func (c *Config) Enablesymbol() *Config {
	c.symbol = true
	return c
}

//Disablesymbol turn off symbols
func (c *Config) Disablesymbol() *Config {
	c.symbol = false
	return c
}

func (c *Config) String() (ret string) {
	SetSeed(time.Now().Unix())
	var available string
	if c.letter {
		available += letter
	}
	if c.number {
		available += number
	}
	if c.symbol {
		available += symbol
	}
	l := len(available)
	for len(ret) < c.length {
		ret += string(available[rand.Intn(l)])
	}
	return ret
}

//New Get a new *Config
func New(length int) *Config {
	return &Config{
		length: length,
		letter: true,
	}
}
