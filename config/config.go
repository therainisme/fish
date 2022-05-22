package config

import "flag"

var BotAddress = flag.String("ba", "", "bot server address, example: http://10.1.1.1:5700")
var FishAddress = flag.String("fa", "", "fish server address, example: http://10.1.1.1:5701")
var FishPort = flag.String("fp", "5701", "fish server port, example: 5701")

func init() {
	flag.Parse()

	if *BotAddress == "" {
		panic("bot address is empty")
	}

	if *FishAddress == "" {
		panic("fish address is empty")
	}
}
