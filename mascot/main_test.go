package mascot_test

import (
	"testing"

	"github.com/tsuccar/devcontainer-gomascot/mascot"
)

func TestMascot(t *testing.T) {

	if mascot.BestMascot() != "Gopher" {
		t.Fatal("wrong mascot :(")
	}

}
