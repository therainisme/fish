package main

import "fish/receiver"

func main() {
	receiver.Listen(":5701")
}
