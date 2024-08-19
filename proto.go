package main

import (
	"bytes"
	"fmt"

	"github.com/tidwall/resp"
	// "io"
	// "log"
	// "github.com/tidwall/resp"
)

const (
	CommandSET = "SET"
	CommandGET = "GET"
	CommandHELLO = "HELLO"
	CommandClient = "Client"
)

type Command interface{

}

type SetCommand struct{
	key,val []byte
}

type HelloCommand struct {
	val string
}

type GetCommand struct{
	key []byte
}

type ClientCommand struct{
	val string
}


func respWriteMap(m map[string]string) []byte {
	buf := &bytes.Buffer{}
	buf.WriteString("%"+fmt.Sprintf("%d\r\n",len(m)))
	rw := resp.NewWriter(buf)
	for k,v := range m{
		rw.WriteString(k)
		rw.WriteString(v)
		buf.WriteString(fmt.Sprintf("%s\r\n",k))
		buf.WriteString(fmt.Sprintf("%s\r\n",v))
	}
	return buf.Bytes()
}