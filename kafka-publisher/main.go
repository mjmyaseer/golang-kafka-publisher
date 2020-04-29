package main

import (
 "Company import path goes here"

)

func main() {
	global.Init()
	defer func() {
		if err := global.Publisher.Close(); err != nil {
			panic(err)
		}
	}()
	interfaces.RouterStart()
}
