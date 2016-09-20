package core_test

import (
	"fmt"
	"github.com/mirrorcomputing/core"
	"testing"
	"time"
)

func TestOuterid_Marshal(t *testing.T) {
	oid := core.Outerid{
		Time: time.Now(),
		Addr: "127.0.0.1:93",
	}
	bs := oid.Marshal()

	newoid, _ := oid.Unmarshal(bs)
	fmt.Println(newoid.String())
}
