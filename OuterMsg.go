package core

import (
	"encoding/binary"
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
	BodyLen uint64
	Body    []byte
}

func (this OuterMsg) Unmarshal(reader io.Reader) (msg OuterMsg, err error) {
	idbs := make([]byte, IdLen)
	idlen, err := reader.Read(idbs)
	if idlen != IdLen || !checkError(err, "Outerid error") {
		return
	}

	Id, err := (Outerid{}).Unmarshal(idbs)
	if !checkError(err, "Outerid error") {
		return
	}
	msg.Id = *Id

	bodylenbs := make([]byte, BodyLenLen)
	bodylenlen, err := reader.Read(bodylenbs)

	if bodylenlen != BodyLenLen || !checkError(err, "bodylen error") {
		return
	}

	bodylen := binary.BigEndian.Uint64(bodylenbs)
	body := make([]byte, bodylen)
	Blen, err := reader.Read(body)
	if Blen != int(bodylen) || !checkError(err, "body error") {
		return
	}
	msg.BodyLen = int(bodylen)
	msg.Body = body
	return
}
