package sio

type Closable interface {
	Close()
}

type ClosableWithError interface {
	Close() error
}

func IsClosable(v interface{}) bool {
	if _, ok := v.(Closable); ok {
		return true
	}
	if _, ok := v.(ClosableWithError); ok {
		return true
	}

	return false
}

func Close(v interface{}) {
	if v, ok := v.(Closable); ok {
		v.Close()
	}

	if v, ok := v.(ClosableWithError); ok {
		_ = v.Close()
	}
}
