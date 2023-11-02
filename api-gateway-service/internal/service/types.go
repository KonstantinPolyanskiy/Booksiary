package service

type ProxyAddrMap struct {
	M map[Root]Handle
}

func NewProxyAddrMap() ProxyAddrMap {
	return ProxyAddrMap{
		M: make(map[Root]Handle),
	}
}
