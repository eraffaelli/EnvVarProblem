package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
)

var (
	RedisHost string
	RedisPort string
)

func show(envvar string, val string) {
	fmt.Println(envvar)
	hexstr := hex.EncodeToString([]byte(val))
	for i := 0; i < len(hexstr); i += 2 {
		bite := hexstr[i : i+2]
		fmt.Printf("%s ", bite)
	}
	fmt.Printf("= %d bytes\n", len(val))
	fmt.Printf("= \"%s\"\n\n", val)
}


func init() {
	if RedisHost = os.Getenv("REDIS_HOST"); RedisHost == "" {
		RedisHost = "localhost"
	}
	if RedisPort = os.Getenv("REDIS_PORT"); RedisPort == "" {
		RedisPort = "6379"
	}

	println("T1")
	temp := RedisHost + `:` + RedisPort
	println(temp)
	println()
	println("T2")
	temp2 := RedisHost + `:` + RedisPort
	println(temp2)
	println()
	println("T3")
	temp3 := RedisHost + RedisPort
	println(temp3)
	println()
	println("Variable :")
	println(RedisHost)
	println(RedisPort)
	println()
	println("Sprintf :")
	println(fmt.Sprintf("RedisHost RedisPort %s:%s\n", RedisHost, RedisPort))
	println(fmt.Sprintf("RedisHost RedisPort %s : %s\n", RedisHost, RedisPort))
	portInt, _ := strconv.Atoi(RedisPort)
	println()
	println("Atoi :\n")
	println(portInt)
	print("From os.getenv\n")
	println(fmt.Sprintf(
		"%s:%s\n",
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT"),
	))
	fmt.Printf("RedisHost : %+v\n", RedisHost)
	println("")
	fmt.Printf("RedisPort : %+v\n", RedisPort)
	fmt.Printf("%T\n", RedisHost)
	fmt.Printf("%T\n", RedisPort)

	fmt.Println("Environment from init:")
	show("REDIS_HOST", os.Getenv("REDIS_HOST"))
	show("REDIS_PORT", os.Getenv("REDIS_PORT"))
}

func main() {
	fmt.Println("Environment from main:")
	show("REDIS_HOST", os.Getenv("REDIS_HOST"))
	show("REDIS_PORT", os.Getenv("REDIS_PORT"))
}