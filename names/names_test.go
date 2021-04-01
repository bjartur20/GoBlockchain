package names_test

import (
	. "github.com/bjartur20/T-419-CADP_OverlayNetworks/names"
	"testing"
	"fmt"
)

func TestMake(t *testing.T) {
	names := Make()
	namesType := fmt.Sprintf("%T", names)
	if namesType != "*names.Names" {
		t.Errorf("have %s want *names.Names", namesType)
	}
}