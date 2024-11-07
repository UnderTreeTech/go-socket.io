package socketio

import (
	"reflect"

	"github.com/UnderTreeTech/go-socket.io/parser"
)

// namespace
const (
	aliasRootNamespace = "/"
	rootNamespace      = ""
)

// message
const (
	clientDisconnectMsg  = "client namespace disconnect"
	enableEncryptEvent   = "shake"
	encryptHandlerSuffix = ":crypto"
)

type readHandler func(c *conn, header parser.Header) error

var (
	defaultHeaderType = []reflect.Type{reflect.TypeOf("")}
)

type cryptoFlag struct{}
