package names

import (
	"errors"
	"time"
)

type Names struct {
	names      map[string]string
	heartbeats map[string]int64
	timeout    int64
}
type Registration struct {
	name    string
	address string // IP:port
}

func (n *Names) GetConnected() (names *map[string]string) {
	names = &n.names
	return
}
func (n *Names) Register(args *Registration) (err error) {
	n.names[args.name] = args.address
	n.heartbeats[args.name] = time.Now().UnixNano()
	return
}
func (n *Names) Unregister(args *string) (err error) {
	delete(n.heartbeats, *args)
	delete(n.names, *args)
	return
}
func (n *Names) Resolve(args *string) (res *string, err error) {
	if val, ok := n.names[*args]; ok {
		res = &val
		return
	}
	err = errors.New("not found")
	return
}
func (n *Names) Heartbeat(args *string) (res int64, err error) {
	if val, ok := n.heartbeats[*args]; ok {
		res = val
		return
	}
	err = errors.New("not found")
	return
}
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
func Make() (res *Names) {
	res = &Names{
		names:      make(map[string]string),
		heartbeats: make(map[string]int64),
		timeout:    int64(time.Second), // 1 second base timeout
	}
	go res.checkHeartbeat()
	return
}
