package names

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestMake(t *testing.T) {
	names := Make()
	namesType := fmt.Sprintf("%T", names)
	if namesType != "*names.Names" {
		t.Errorf("have %s want *names.Names", namesType)
	}
}
func TestRegister(t *testing.T) {
	names := Make()
	node := Registration{"A", "123"}
	err := names.Register(&node)
	if err != nil {
		t.Errorf("Register failed")
	}
	if address, ok := names.names[node.name]; ok {
		if address != node.address {
			t.Errorf("Did not match input")
		}
	} else {
		t.Error("Was not put into names map")
	}
	if hb, ok := names.heartbeats[node.name]; ok {
		now := time.Now().UnixNano()
		if hb < now-5*int64(math.Pow10(8)) || hb > now {
			t.Error("Did not match time")
		}
	} else {
		t.Error("Was not put into heartbeat map")
	}
}
func TestUnregister(t *testing.T) {
	names := Make()
	A := Registration{"A", "123"}
	names.Register(&A)
	B := Registration{"B", "124"}
	names.Register(&B)
	names.Unregister(&A.name)
	if len(names.names) != 1 {
		t.Error("Name was not deleted from names map")
	}
	if len(names.heartbeats) != 1 {
		t.Error("Name was not deleted from heartbeats map")
	}
}
func TestResolve(t *testing.T) {
	names := Make()
	A := Registration{"A", "123"}
	names.Register(&A)
	B := Registration{"B", "124"}
	names.Register(&B)
	// Test find it
	address, err := names.Resolve(&A.name)
	if err != nil {
		t.Errorf("Resolve failed: %s", err)
	}
	if *address != A.address {
		t.Error("Did not find correct address")
	}
	// Test don't find it
	name := "C"
	address, err = names.Resolve(&name)
	if err == nil {
		t.Error("Resolve should have failed")
	} else {
		if address != nil {
			t.Error("should not have returned an address")
		}
	}
}
func TestHeartbeat(t *testing.T) {
	names := Make()
	A := Registration{"A", "123"}
	names.Register(&A)
	B := Registration{"B", "124"}
	names.Register(&B)
	// Test find it
	_, err := names.Heartbeat(&A.name)
	if err != nil {
		t.Errorf("Heartbeat failed: %s", err)
	}
	// Test don't find it
	name := "C"
	_, err = names.Heartbeat(&name)
	if err == nil {
		t.Error("Heartbeat should have failed")
	}
}
func TestCheckHeartbeat(t *testing.T) {
	names := Make()
	A := Registration{"Dead heartbeat", "123"}
	names.Register(&A)
	names.heartbeats[A.name] = time.Now().UnixNano() - 120*int64(math.Pow10(9))

	time.Sleep(time.Second)
	if len(names.names) != 0 || len(names.heartbeats) != 0 {
		t.Error("node should have unregistered")
	}
}
