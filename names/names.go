package names

// import (
// 	"time"
// )

type Names struct {
	names map[string]string
	heartbeat map[string]int64
	timeout int64
}

type Registration struct {
	name string
	address string // IP:port
}

func (*Names) Register(args *Registration, res *string) error {
	//n := Names{}
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

func (names *Names) checkHeartbeat() {
	return
	// for {
		// time.Sleep(names.timeout / 2)
		// Remove all entries whose heartbeat is older than 60 seconds
	// }
}

func  Make() (res *Names) {
	res = &Names{ }
	go res.checkHeartbeat()
	return
}
