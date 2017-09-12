package libs


import (
	"encoding/json"
	"errors"
	"net"
)

var (
	emptyProto    = Proto{}
	emptyJSONBody = []byte("{}")

	ErrProtoPackLen   = errors.New("default server codec pack length error")
	ErrProtoHeaderLen = errors.New("default server codec header length error")
)


// for tcp
const (
	MaxBodySize = int32(1 << 10)
)

const (
	// size
	PackSize      = 4
	HeaderSize    = 2
	VerSize       = 2
	OperationSize = 4
	SeqIdSize     = 4
	RawHeaderSize = PackSize + HeaderSize + VerSize + OperationSize + SeqIdSize
	MaxPackSize   = MaxBodySize + int32(RawHeaderSize)
	// offset
	PackOffset      = 0
	HeaderOffset    = PackOffset + PackSize
	VerOffset       = HeaderOffset + HeaderSize
	OperationOffset = VerOffset + VerSize
	SeqIdOffset     = OperationOffset + OperationSize
)

type Proto struct {
	Ver int16				`json:"ver"`
	Operation int32			`json:"op"`
	SeqId int32				`json:"seq"`
	Body json.RawMessage	`json:"body"`
}

func (p *Proto) ReadTcp(r *net.TCPConn) (err error){
	var (
		bodyLen   int
		headerLen int16
		packLen   int32
	)

	head := [RawHeaderSize]byte{}
	if _,err := r.Read(head[0:]);err != nil{
		return err
	}
	packLen = BigEndian.Int32(head[PackOffset:HeaderOffset])
	headerLen = BigEndian.Int16(head[HeaderOffset:VerOffset])
	p.Ver = BigEndian.Int16(head[VerOffset:OperationOffset])
	p.Operation = BigEndian.Int32(head[OperationOffset:SeqIdOffset])
	p.SeqId = BigEndian.Int32(head[SeqIdOffset:])
	if packLen > MaxPackSize{
		return ErrProtoPackLen
	}
	if headerLen != RawHeaderSize{
		return ErrProtoHeaderLen
	}
	if bodyLen = int(packLen - int32(headerLen));bodyLen > 0{
		buf := make([]byte,bodyLen)
		r.Read(buf)
		p.Body = buf
	}else{
		p.Body = nil
	}
	return
}

func (p *Proto) WriteTcp(w *net.TCPConn) (err error){
	var (
		buf     []byte
		packLen int32
	)

	packLen = RawHeaderSize + int32(len(p.Body))
	buf = make([]byte,packLen)
	BigEndian.PutInt32(buf[PackOffset:], packLen)
	BigEndian.PutInt16(buf[HeaderOffset:], int16(RawHeaderSize))
	BigEndian.PutInt16(buf[VerOffset:], p.Ver)
	BigEndian.PutInt32(buf[OperationOffset:], p.Operation)
	BigEndian.PutInt32(buf[SeqIdOffset:], p.SeqId)
	copy(buf[RawHeaderSize:],p.Body)
	w.Write(buf)
	return
}
