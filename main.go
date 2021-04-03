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
	"net/http"
	"os"
	"time"

	"github.com/bjartur20/T-419-CADP_OverlayNetworks/names"
	"github.com/nictuku/dht"
)

const (
	httpPortTCP = 8711
	numTarget   = 10
	exampleIH   = "94a315e2cf8015b2f635d79aab592e6db557d5ea"
)

// drainresults loops, printing the address of nodes it has found.
func drainresults(d *dht.DHT) {
	fmt.Println("========================= Naming Service =========================")
	fmt.Println("Note that there are many bad nodes that reply to anything you ask.")
	fmt.Println("Peers found:")
	for r := range d.PeersRequestResults {
		for _, peers := range r { // Returns our info hash and list of map of peers
			for count, x := range peers { // index (number) i think and Encoded Peer Address
				fmt.Printf("  %+v: %+v\n", count, dht.DecodePeerAddress(x))
				if count >= 10 {
					return
				}
				count++
			}
		}
	}
}

func main() {
	nameService := names.Make()
	fmt.Printf("%+v\n", nameService)
	flag.Parse()
	 
	ih, _ := dht.DecodeInfoHash("45b3d693cff285975f622acaeb75c5626acaff6f")
	if len(flag.Args()) == 1 {
		var err error
		ih, err = dht.DecodeInfoHash(flag.Args()[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "DecodeInfoHash error: %v\n", err)
			os.Exit(1)
		}
	}

	// Starts a DHT node with the default options. It picks a random UDP port. To change this, see dht.NewConfig.
	d, err := dht.New(nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "New DHT error: %v", err)
		os.Exit(1)
	}

	// For debugging.
	go http.ListenAndServe(fmt.Sprintf(":%d", httpPortTCP), nil)

	if err = d.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "DHT start error: %v", err)
		os.Exit(1)
	}

	go drainresults(d)

	for {
		d.PeersRequest(string(ih), false)
		time.Sleep(5 * time.Second)
	}
}
