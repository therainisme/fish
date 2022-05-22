package main

import (
	"fish/config"
	"fish/receiver"
)

func main() {
	receiver.Listen(":" + *config.FishPort)
}
