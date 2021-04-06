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
	"flag"
	"fmt"
	"net"
	"net/url"
	"net/http"
	"time"

	"github.com/bjartur20/T-419-CADP_OverlayNetworks/names"
	"github.com/bjartur20/T-419-CADP_OverlayNetworks/debug_logger"
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

func drainResults(d *dht.DHT, n *names.Names) {
	for r := range d.PeersRequestResults {
		for _, peers := range r {
			for _, x := range peers {
				address := dht.DecodePeerAddress(x)
				hostname, err := getHostname(&address)
				if err != nil {
					continue
				}
				fmt.Printf("Peer connected: %v (%v)\n", *hostname, address)
				if n != nil {
					node := names.MakeRegistration(hostname, &address)
					n.Register(node)
				}
			}
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
	node.DebugLogger = &debug_logger

	return node, nil
}

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "1.1.1.1:80")
	if err != nil {
		fmt.Errorf("%v", err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func main() {
	flag.Parse()
	infoHash, _ := dht.DecodeInfoHash("d1c5676ae7ac98e8b19f63565905105e3c4c37a2")
	var nameService *names.Names
	var node *dht.DHT

	if len(flag.Args()) == 0 {
		fmt.Printf("No router argument, starting router node...\n")

		// Initialize the naming service
		nameService = names.Make()

		// Start route node
		node, _ = startNode("", string(infoHash))
		name := GetOutboundIP()
		ip := fmt.Sprintf("%s:%d", name, node.Port())
		go http.ListenAndServe(":8711", nil)
		fmt.Printf("Please connect to this router: %v. With info hash: %v\n", ip, infoHash)
	} else {
		nameService = nil
		node, _ = startNode(flag.Args()[0], string(infoHash))
		name := GetOutboundIP()
		ip := fmt.Sprintf("%s:%d", name, node.Port())
		go http.ListenAndServe(":8712", nil)
		fmt.Printf("Started node: %v. With info hash: %v\n", ip, infoHash)
	}

	// go http.ListenAndServe(fmt.Sprintf(":%d", httpPortTCP), nil)

	go drainResults(node, nameService)

	for {
		node.PeersRequest(string(infoHash), true)
		time.Sleep(5 * time.Second)
	}
}
