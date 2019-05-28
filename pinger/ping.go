package pinger

import (
	"context"
	"runtime"
	"time"

	"github.com/sparrc/go-ping"
)

type Peer struct {
	Address  string
	Status   bool
	PingTime time.Duration
	Ctx      *context.Context
}

func Ping(address string, ctx *context.Context, peer *Peer) {
	pinger, err := ping.NewPinger(address)
	if err != nil {
		panic(err)
	}
	if runtime.GOOS == "windows" {
		pinger.SetPrivileged(true)
	}

	pinger.Count = 1
	pinger.OnRecv = func(pkt *ping.Packet) {
		// log.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
		// 	pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
		if pkt.Nbytes > 0 {
			peer.Status = true
			peer.PingTime = pkt.Rtt
			// pinger.Stop()
			return
		} else {
			// pinger.Stop()
			return
		}

	}

	// pinger.OnFinish = func(stats *ping.Statistics) {
	// 	// log.Printf("\n--- %s ping statistics ---\n", stats.Addr)
	// 	// log.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
	// 	// 	stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
	// 	// log.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
	// 	// 	stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
	// }

	// log.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	pinger.Run()
}
