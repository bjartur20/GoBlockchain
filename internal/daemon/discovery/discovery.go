package discovery

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/nictuku/dht"
	"github.com/bjartur20/GoBlockchain/internal/daemon/debug_logger"
)

const (
	GenesisBlockHash = "a9f5ad541c1afbbe19a617a9ca28b3826b66db5e"
	anounceOurselves = false
)

func drainNodes(n *dht.DHT) {
	for r := range n.PeersRequestResults {
		for _, peers := range r {
			for _, x := range peers {
				ip, _, err := net.SplitHostPort(dht.DecodePeerAddress(x))
				if err != nil {
					log.Println(err)
					continue
				}
				ptr, err := net.LookupAddr(ip)
				if err != nil {
					log.Println(err)
					continue
				}
				log.Println("peer:", ptr[0])
			}
		}
	}
}

func createDHTConfig(routers string) (c *dht.Config) {
	c = dht.NewConfig()
	c.SaveRoutingTable = false
	c.DHTRouters = routers
	c.Port = 0 // Sets random availabel port

	return
}

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
    conn, err := net.Dial("udp", "1.1.1.1:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP
}

func Run(routers string) {
	ih, err := dht.DecodeInfoHash(GenesisBlockHash)

	// Create config
	c := createDHTConfig(routers)

	// Start a DHT node on random port.
	d, err := dht.New(c)
	if err != nil {
		log.Fatal(err)
	}
	err = d.Start()
	if err != nil {
		log.Fatal(err)
	}

	// Start debug logger
	d.DebugLogger = &debug_logger.DebugLogger{}

	// Print node's name+ip
	name := GetOutboundIP()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%s:%d\n", name, d.Port())

	// Drain discovered nodes in a seperate goroutine.
	go drainNodes(d)

	// Keep requesting for more pairs in an endless loop.
	for {
		d.PeersRequest(string(ih), anounceOurselves)
		time.Sleep(5 * time.Second)
	}
}
