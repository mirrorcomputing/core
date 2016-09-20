package core

import (
	"time"
)

/*
外世界ID

30 byte:time
20 byte:addr
*/
type Outerid struct  {
	time time.Time
	addr string
}

func Unmarshal(bs []byte)(self *Outerid,err error){
	self.time,err=time.Parse("2009-11-10 23:00:00",string(bs[:30]))
	self.addr=string(bs[30:])
	return self,err
}
func (this *Outerid)Marshal()(bs []byte){
	ts:=[]byte(this.time.String()[:19])
	ads:=[]byte(this.addr)
	bs=make([]byte,50)
	copy(bs[:30],ts)
	copy(bs[30:],ads)
	return bs
}