package core

import (
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
	self.Time, err = time.Parse("2009-11-10 23:00:00", string(bs[:30]))
	self.Addr = string(bs[30:])
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
func (this *Outerid) String() (s string) {
	s = string(this.Marshal())
	return
}
