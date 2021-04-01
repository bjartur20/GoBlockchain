package names

import (
	"time"
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

func (*Names) Unregister(args *string, res *string) error {
	return nil
}

func (*Names) Resolve(args *string, res *string) error {
	return nil
}

func (*Names) Heartbeat(args *string, res *string) error {
	return nil
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
	res = &Names{}
	res.timeout = 42 // Some base timeout
	go res.checkHeartbeat()
	return
}
