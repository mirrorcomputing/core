package core

import (
	"encoding/binary"
	"fmt"
	"io"
)

const (
	IdLen      = OuteridLen
	BodyLenLen = 8
)

/*
OuterMsg 报文格式

(2016/9/1 9:21:20,192.127.21.12:23423)
bodylen:
body...
*/
type OuterMsg struct {
	Id      Outerid
	GoId    Outerid
	BodyLen uint64
	Body    []byte
}

func (this *OuterMsg) WriteString(s string) *OuterMsg {
	this.Body = []byte(s)
	this.BodyLen = uint64(len([]byte(s)))
	return this
}
func (this *OuterMsg) Write(bs []byte) *OuterMsg {
	this.Body = bs
	this.BodyLen = uint64(len(bs))
	return this
}

// region 序列化
func (this *OuterMsg) Marshal() (bs []byte) {
	bs = make([]byte, OuteridLen*2+8+int(this.BodyLen))
	copy(bs[:OuteridLen], this.Id.Marshal())
	copy(bs[OuteridLen:OuteridLen*2], this.GoId.Marshal())
	binary.BigEndian.PutUint64(bs[OuteridLen*2:OuteridLen*2+8], this.BodyLen)
	copy(bs[OuteridLen*2+8:], this.Body)
	return
}
func (this OuterMsg) Read(reader io.Reader) (msg OuterMsg, err error) {
	fmt.Println("1")
	idbs := make([]byte, IdLen)
	idlen, err := reader.Read(idbs)
	if idlen != IdLen || !checkError(err, "Outerid error") {
		return
	}
	fmt.Println("2")
	Id, err := (Outerid{}).Unmarshal(idbs)
	if !checkError(err, "Outerid error") {
		return
	}
	fmt.Println("3")
	msg.Id = *Id
	goidbs := make([]byte, IdLen)
	goidlen, err := reader.Read(goidbs)
	if goidlen != IdLen || !checkError(err, "Outerid error") {
		return
	}
	fmt.Println("4")
	GoId, err := (Outerid{}).Unmarshal(goidbs)
	fmt.Println(len(goidbs), GoId)
	if !checkError(err, "Outerid error") {
		return
	}
	fmt.Println("5")
	msg.GoId = *GoId
	bodylenbs := make([]byte, BodyLenLen)
	bodylenlen, err := reader.Read(bodylenbs)
	fmt.Println("6")
	if bodylenlen != BodyLenLen || !checkError(err, "bodylen error") {
		return
	}
	fmt.Println("7")
	bodylen := binary.BigEndian.Uint64(bodylenbs)
	body := make([]byte, bodylen)
	Blen, err := reader.Read(body)
	if Blen != int(bodylen) || !checkError(err, "body error") {
		return
	}
	fmt.Println("8")
	msg.BodyLen = bodylen
	msg.Body = body
	return
}
func (this OuterMsg) Unmarshal(bs []byte) (msg OuterMsg, err error) {
	id := &Outerid{}
	goid := &Outerid{}
	id, err = (Outerid{}).Unmarshal(bs[:IdLen])
	goid, err = (Outerid{}).Unmarshal(bs[IdLen : IdLen*2])
	msg.Id = *id
	msg.GoId = *goid
	msg.BodyLen = binary.BigEndian.Uint64(bs[IdLen*2 : IdLen*2+BodyLenLen])
	msg.Body = bs[IdLen*2+BodyLenLen:]
	//
	return
}

// endregion
