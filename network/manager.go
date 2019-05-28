package network

import (
	"context"
	"log"
)

func InitManager() {
	log.Println("Initializing the scan manager")

	// count, IPrange := CalculateCIDR("132.186.255.103/24")

	// statusChannels := make([]chan string, count)

	// log.Println(count)

	// for _, c := range IPrange {
	// 	go Ping(c, context.Background())
	// }

}

var placing = make(map[context.Context]int)

func SetPlace(ctx context.Context) {

}

func GetPlace(ctx context.Context) {

}
