// Package names implements the name service to map logical node names to
// host names and port numbers, needed to bootstrap an overlay network.	
package names

// Imports errors for error handling and time for heartbeat.
import (
	"errors"
	"time"
)

// Type for storing names. Includes a map of names and addresses,
// a map of names and heartbeats, and a timeout cap.
type Names struct {
	names      map[string]string
	heartbeats map[string]int64
	timeout    int64
}

// Type for registration data, name and address pair.
type Registration struct {
	name    string
	address string // IP:port
}

// Returns all names on the service.
func (n *Names) GetConnected() (names *map[string]string) {
	names = &n.names
	return
}

// Registers a client to the service using a Registration.
func (n *Names) Register(args *Registration) (err error) {
	n.names[args.name] = args.address
	n.heartbeats[args.name] = time.Now().UnixNano()
	return
}

// Unregisters a client from the service.
func (n *Names) Unregister(args *string) (err error) {
	delete(n.heartbeats, *args)
	delete(n.names, *args)
	return
}

// Resloves a hostname to an address.
func (n *Names) Resolve(args *string) (res *string, err error) {
	if val, ok := n.names[*args]; ok {
		res = &val
		return
	}
	err = errors.New("not found")
	return
}

// Returns the heartbeat of all clients that are registered.
func (n *Names) Heartbeat(args *string) (res int64, err error) {
	if val, ok := n.heartbeats[*args]; ok {
		res = val
		return
	}
	err = errors.New("not found")
	return
}

// Checks heartbeat of all addresses in registration.
// Unregisters unresponsive clients.
func (n *Names) checkHeartbeat() {
	for {
		time.Sleep(time.Duration(n.timeout))
		timeNow := time.Now().UnixNano()
		// Remove all entries whose heartbeat is older than 60 seconds
		for host, heartbeat := range n.heartbeats {
			if timeNow-heartbeat > int64(60*time.Second) {
				n.Unregister(&host)
			}
		}
	}
}

// Registers a new Registration to the naming service.
func MakeRegistration(name, address *string) (res *Registration) {
	res = &Registration{
		name: *name,
		address: *address,
	}
	return
}

// Creates a new naming service.
func Make() (res *Names) {
	res = &Names{
		names:      make(map[string]string),
		heartbeats: make(map[string]int64),
		timeout:    int64(time.Second), // 1 second base timeout
	}
	go res.checkHeartbeat()
	return
}
