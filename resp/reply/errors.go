package reply

// 错误响应 位置错误、参数数量错误、语法错误、类型错误、协议错误

// UnknownErrReply 未知错误
type UnknownErrReply struct {
}

var unknownErrBytes = []byte("-Err unknown\r\n")

func (r *UnknownErrReply) ToBytes() []byte {
	return unknownErrBytes
}

func (r *UnknownErrReply) Error() string {
	return "Err unknown"
}

// ArgNumErrReply 参数数量不对错误提示
type ArgNumErrReply struct {
	Cmd string
}

func (r *ArgNumErrReply) ToBytes() []byte {
	return []byte("_ERR wrong number of arguments for '" + r.Cmd + "' command")
}

func (r *ArgNumErrReply) Error() string {
	return "ERR wrong number of arguments for '" + r.Cmd + "' command"
}

func MakeArgNumErrReply(cmd string) *ArgNumErrReply {
	return &ArgNumErrReply{
		Cmd: cmd,
	}
}

// SyntaxErrReply 语法错误
type SyntaxErrReply struct {
}

var syntaxErrBytes = []byte("-Err syntax err\r\n")
var theSyntaxErrReply = &SyntaxErrReply{}

// ToBytes marshals redis.Reply
func (r *SyntaxErrReply) ToBytes() []byte {
	return syntaxErrBytes
}

func (r *SyntaxErrReply) Error() string {
	return "Err syntax error"
}

// WrongTypeErrReply 错误的类型
type WrongTypeErrReply struct {
}

var wrongTypeErrBytes = []byte("-WRONGTYPE Operation against a key holding the wrong kind of value\r\n")

func (r *WrongTypeErrReply) ToBytes() []byte {
	return wrongTypeErrBytes
}

func (r *WrongTypeErrReply) Error() string {
	return "-WRONGTYPE Operation against a key holding the wrong kind of value\r\n"
}

// ProtocolErrReply 协议错误提示
type ProtocolErrReply struct {
	Msg string
}

func (r *ProtocolErrReply) ToBytes() []byte {
	return []byte("-ERR Protocol error: '" + r.Msg + "'\r\n")
}

func (r *ProtocolErrReply) Error() string {
	return "ERR Protocol error: '" + r.Msg
}
