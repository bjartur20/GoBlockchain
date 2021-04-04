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

	"github.com/bjartur20/T-419-CADP_OverlayNetworks/names"
	"github.com/nictuku/dht"
)

// const (
// 	httpPortTCP = 8711
// 	numTarget   = 10
// 	exampleIH   = "94a315e2cf8015b2f635d79aab592e6db557d5ea"
// )

func getHostname(addess *string, i int) (res *string) {
	res = &[]string{"Emma", "Olivia", "Ava", "Isabella", "Sophia", "Charlotte", "Mia", "Amelia", "Harper", "Evelyn"}[i%10]
	return
}

func processIncommingRequests(d *dht.DHT, n *names.Names) {
	count := 0
	for r := range d.PeersRequestResults { // Collect addresses requested by PeersRequest() Call
		for _, peers := range r { // Returns our info hash and list of map of peers
			for _, encodedAddress := range peers { // index Encoded Peer Address
				address := dht.DecodePeerAddress(encodedAddress)
				name := *getHostname(&address, count)
				node := names.MakeRegistration(&name, &address)
				n.Register(node)
				fmt.Printf("%s: %s\n", name, address)
				count++
			}
		}
	}
}

func drainResults(n *dht.DHT, ih string) error {
	count := 1
	for {
		select {
		case r := <-n.PeersRequestResults:
			for _, peers := range r {
				for _, x := range peers {
					address := dht.DecodePeerAddress(x)
					url, _ := url.Parse("http://" + address)
					name, _ := net.LookupAddr(url.Hostname())
					fmt.Printf("Hostname: %v\n", name[0])
					count++
				}
			}
		case <-time.Tick(time.Second / 5):
			n.PeersRequest(ih, true)
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
	// nameService := names.Make()

	infoHash, _ := dht.DecodeInfoHash("d1c5676ae7ac98e8b19f63565905105e3c4c37a2")

	routerNode, _ := startNode("", string(infoHash))
	name := GetOutboundIP()
	router := fmt.Sprintf("%s:%d", name, routerNode.Port())
	fmt.Printf("Please connect to this router: %v. With info hash: %v\n", router, infoHash)

	startNode(router, string(infoHash))
	startNode(router, string(infoHash))
	startNode(router, string(infoHash))
	startNode(router, string(infoHash))

	// go drainResults(n1, string(infoHash))
	// go drainResults(n2, string(infoHash))
	// go drainResults(n3, string(infoHash))
	// go drainResults(n4, string(infoHash))

	go drainResults(routerNode, string(infoHash))
	// go drainResults(n3, string(infoHash), 3, 1000*time.Second)

	// go processIncommingRequests(d, nameService)

	for {
		time.Sleep(100 * time.Second)
	}
}
