package tcpclient

import (
	"sync"
)

var ConnChanMap *TcpConnChanMap
type TcpConnChannel chan string

type TcpConnChanMap struct {
	connChanMap map[string]TcpConnChannel
	rwMutex sync.RWMutex
}

func CreateTcpConnChanMap() {
	*ConnChanMap = TcpConnChanMap {
		connChanMap : make(map[string]TcpConnChannel),
	}
}