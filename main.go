// Runs a node on a random UDP port that attempts to collect 10 peers for an
// infohash, then keeps running as a passive DHT node.
//
// IMPORTANT: if the UDP port is not reachable from the public internet, you
// may see very few results.
//
// To collect 10 peers, it usually has to contact some 1k nodes. It's much easier
// to find peers for popular infohashes. This process is not instant and should
// take a minute or two, depending on your network connection.
//
//
// There is a builtin web server that can be used to collect debugging stats
// from http://localhost:8711/debug/vars.
package main

import (
	"fmt"
	"net"
	"net/url"
	"time"
	"flag"

	"github.com/bjartur20/T-419-CADP_OverlayNetworks/names"
	"github.com/nictuku/dht"
)

func getHostname(address *string) (hostname *string, err error) {
	url, err := url.Parse("http://" + *address)
	if err != nil {
		return
	}
	hostnames, err := net.LookupAddr(url.Hostname())
	if err != nil {
		return
	}
	hostname = &hostnames[0]
	return
}

func drainResults(d *dht.DHT, n *names.Names, ih string) error {
	count := 1
	for {
		select {
		case r := <-d.PeersRequestResults:
			for _, peers := range r {
				for _, x := range peers {
					address := dht.DecodePeerAddress(x)
					hostname, err := getHostname(&address)
					if err != nil {
						continue
					}
					fmt.Printf("Peer connected: %v (%v)\n", *hostname, address)
					node := names.MakeRegistration(hostname, &address)
					n.Register(node)
					count++
				}
			}
		case <-time.Tick(time.Second / 5):
			d.PeersRequest(ih, true)
		}
	}
}

func startNode(routers string, ih string) (*dht.DHT, error) {
	c := dht.NewConfig()
	c.SaveRoutingTable = false
	c.DHTRouters = routers
	c.Port = 0
	node, err := dht.New(c)
	if err != nil {
		return nil, err
	}
	if err = node.Start(); err != nil {
		return nil, err
	}
	node.PeersRequest(ih, true)
	return node, nil
}

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Errorf("%v", err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func main() {
	infoHash, _ := dht.DecodeInfoHash("d1c5676ae7ac98e8b19f63565905105e3c4c37a2")
	
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("No router argument, starting router node...\n")

		// Initialize the naming service
		nameService := names.Make()

		// Start route node
		routerNode, _ := startNode("", string(infoHash))
		name := GetOutboundIP()
		router := fmt.Sprintf("%s:%d", name, routerNode.Port())
		fmt.Printf("Please connect to this router: %v. With info hash: %v\n", router, infoHash)
		go drainResults(routerNode, nameService, string(infoHash))
	} else {
		startNode(flag.Args()[0], string(infoHash))
	}

	// startNode(router, string(infoHash))
	// startNode(router, string(infoHash))
	// startNode(router, string(infoHash))


	for {
		time.Sleep(100 * time.Second)
	}
}
