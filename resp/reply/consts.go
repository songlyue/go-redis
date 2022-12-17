package reply

// 处理了pong响应、ok响应、空块响应、空的list响应、‘什么都没有’的响应

// PongReply 给ping响应一个pong
type PongReply struct{}

var pongBytes = []byte("+PONG\r\n")

// ToBytes 响应一个ping字节数组
func (r *PongReply) ToBytes() []byte {
	return pongBytes
}

// OkReply OK的响应
type OkReply struct {
}

var okBytes = []byte("+OK\r\n")

// 返回ok的字节数组
func (r *OkReply) ToBytes() []byte {
	return okBytes
}

var theOkReply = new(OkReply)

func MakeOKReply() *OkReply {
	return theOkReply
}

var nullBulkBytes = []byte("$-1\r\n")

type NullBulkReply struct {
}

func (r *NullBulkReply) ToBytes() []byte {
	return nullBulkBytes
}

func MakeNullBulkReply() *NullBulkReply {
	return &NullBulkReply{}
}

var emptyMultiBulkBytes = []byte("*0\r\n")

type EmptyMultiBulkReply struct {
}

// ToBytes marshal redis.Reply
func (r *EmptyMultiBulkReply) ToBytes() []byte {
	return emptyMultiBulkBytes
}

// NoReply respond nothing, for commands like subscribe
type NoReply struct{}

var noBytes = []byte("")

// ToBytes marshal redis.Reply
func (r *NoReply) ToBytes() []byte {
	return noBytes
}
