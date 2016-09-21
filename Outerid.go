package core

import (
	"fmt"
	"github.com/mirrorcomputing/core/bytes"
	"time"
)

const OuteridLen = 50

/*
外世界ID

30 byte:time
20 byte:addr
*/
type Outerid struct {
	Time time.Time
	Addr string
}

func (this Outerid) Unmarshal(bs []byte) (self *Outerid, err error) {

	self = &Outerid{}
	t := string(bytes.Trim0(bs[:30]))

	withNanos := "2006-01-02 15:04:05"

	self.Time, err = time.Parse(withNanos, t)

	self.Addr = string(bytes.Trim0(bs[30:]))
	fmt.Println("outerid", t)
	return self, err
}
func (this *Outerid) Marshal() (bs []byte) {
	ts := []byte(this.Time.String()[:19])
	ads := []byte(this.Addr)
	bs = make([]byte, 50)
	copy(bs[:30], ts)
	copy(bs[30:], ads)
	return bs
}
func (this Outerid) String() string {
	s := this.Time.String()[:19] + "," + this.Addr
	return s
}
