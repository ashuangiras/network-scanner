package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/ashuangiras/network-scanner/network"

	"github.com/ashuangiras/network-scanner/pinger"
)

func main() {
	log.Println("==================== Network Scanner ==================== ")

	cidr := flag.String("cidr", "192.168.1.1/24", "CIDR block to scan")
	flag.Parse()

	peerConn := make(map[string]*pinger.Peer)

	_, IPrange := network.CalculateCIDR(*cidr)

	for _, ip := range IPrange {
		ctx := context.Background()
		peerConn[ip] = &pinger.Peer{
			Address: ip,
			Ctx:     &ctx,
		}
		go pinger.Ping(ip, peerConn[ip].Ctx, peerConn[ip])

	}

	log.Println("Scanning all the peers . . .")
	time.Sleep(time.Second * 5)

	log.Println("Total Peers Online : ", len(peerConn))
	for _, peer := range peerConn {
		if peer.Status {
			log.Printf("  : %6s - Online - %6v\t\n", peer.Address, peer.PingTime)
		}
	}
	// }

}
