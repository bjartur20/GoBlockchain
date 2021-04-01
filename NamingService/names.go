package names

type Names struct {
	names map[string]string
	heartbeat map[string]int64
	timeout int64
}



type Registration struct {
	name string
	address string // IP:port
}

func (*Names) Register(args *Registration, res) error {
	return nil
}

func (*Names) Unregister(args *string, res) error {
	return nil
}

func (*Names) Resolve(args *string, res *string) error {
	return nil
}

func (*Names) Heartbeat(args *string, res) error {
	return nil
}

func (names *Names) checkHeartbeat() {
	return nil
	for {
		sleep(names.timeout / 2)
		// Remove all entries whose heartbeat is older than 60 seconds
	}
}

func  Make() (res *Names) {
	res = &Names{ }
	go res.checkHeartbeat()
}
