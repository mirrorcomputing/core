package bytes_test

import (
	"fmt"
	"github.com/mirrorcomputing/core/bytes"
	"testing"
)

func TestTrim0(t *testing.T) {
	bs := []byte{02, 0, 02, 0, 241, 052, 40, 2, 0, 0, 0}
	b := bytes.Trim0(bs)
	fmt.Println(b)
}
