package core_test

import (
	"fmt"
	"github.com/mirrorcomputing/core"
	"testing"
	"time"
)

func TestOuterMsg_Unmarshal(t *testing.T) {
	oid := core.Outerid{
		Time: time.Now(),
		Addr: "127.0.0.1:93",
	}
	omsg := core.OuterMsg{
		Id:      oid,
		BodyLen: 5,
		Body:    []byte{1, 2, 3, 4, 5},
	}

	bs := omsg.Marshal()

	newoid, err := omsg.Unmarshal(bs)
	fmt.Println(newoid.Body, err)
}
