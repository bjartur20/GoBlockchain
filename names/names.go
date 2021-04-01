package names

import (
	"time"
	"errors"
)

type Names struct {
	names map[string]string
	heartbeat map[string]int64
	timeout int64
}

type Registration struct {
	name string
	address string // IP:port
}

func (n *Names) Register(args *Registration, res *string) error {
	// Not sure if any of the things I am doing here is correct
	n.names[args.name] = args.address
	n.heartbeat[args.name] = time.Now().UnixNano() // This is converting time to int64
	return nil
}

func (n *Names) Unregister(args *string, res *string) error {
	// Remove host from heartbeat and from names
	delete(n.heartbeat, *args)
	delete(n.names, *args)
	return nil
}

func (n *Names) Resolve(args *string) (*string, error) {
	if val, ok := n.names[*args]; ok {
		return &val, nil
	}
	return nil, errors.New("not found")
}

func (n *Names) Heartbeat(args *string) (*int64, error) {
	if val, ok := n.heartbeat[*args]; ok {
		return &val, nil
	}
	return nil, errors.New("not found")
}

func (n *Names) checkHeartbeat() {
	for {
		time.Sleep(time.Duration(n.timeout))
		timeNow := time.Now().UnixNano()
		// Remove all entries whose heartbeat is older than 60 seconds
		for host, heartbeat := range n.heartbeat {
			if heartbeat - timeNow > 60 { // Im not sure if I am using the right unit here
				res := "i dont know what this is for"
				n.Unregister(&host, &res)
			}
		}
	}
}

func Make() (res *Names) {
	res = &Names{
		names:     make(map[string]string),
        heartbeat: make(map[string]int64),
        timeout:   42, // Some base timeout 
	}
	go res.checkHeartbeat()
	return
}
