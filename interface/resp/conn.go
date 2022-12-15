package resp

type Connection interface {
	Write([]byte) error
	GetDbIndex() int
	SelectDB(int)
}
