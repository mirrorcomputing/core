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
	fmt.Println(len(bs))
	fmt.Println(bs)
	newoid, err := oid.Unmarshal(bs[:19])
	fmt.Println(err, newoid.String())
}
