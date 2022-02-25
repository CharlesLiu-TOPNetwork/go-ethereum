package common

import "fmt"

var (
	ErrRpcReturnError   = fmt.Errorf("rpc return error")
	ErrRpcValueNotfound = fmt.Errorf("rpc value not found")
)
