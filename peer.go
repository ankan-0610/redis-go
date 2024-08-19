package main

import (
	// "fmt"
	// "log/slog"
	// "fmt"
	"io"
	"log"
	"net"

	"github.com/tidwall/resp"
)

type Peer struct{
	conn net.Conn
	msgCh chan Message
	delCh chan *Peer
}

func (p *Peer) Send(msg []byte) (int,error){
	return p.conn.Write(msg)
}

func NewPeer(conn net.Conn, msgCh chan Message, delCh chan *Peer) *Peer{
	return &Peer{
		conn:conn,
		msgCh: msgCh,
		delCh: delCh,
	}
}

func (p *Peer) readLoop() error{
	rd := resp.NewReader(p.conn)
	for {
		v, _, err := rd.ReadValue()
		if err == io.EOF {
			p.delCh <- p
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("Read %s\n", v.Type())
		var cmd Command
		if v.Type() == resp.Array {
			rawCmd := v.Array()[0]
			switch rawCmd.String(){
			case CommandGET:
				cmd = GetCommand{
					key: v.Array()[1].Bytes(),
				}
			case CommandSET:
				cmd = SetCommand{
					key: v.Array()[1].Bytes(),
					val: v.Array()[2].Bytes(),
				}
			case CommandHELLO:
				cmd = HelloCommand{
					val: v.Array()[1].String(),
				}
			case CommandClient:
				cmd = ClientCommand{
					val: v.Array()[1].String(),
				}
			}
			p.msgCh <- Message{
				cmd: cmd,
				peer: p,
			}
			// for _, value := range v.Array() {
				
				// switch v.String(){
				// case CommandSET:
				// 	if len(v.Array())!=3{
				// 		return fmt.Errorf("invalid no. of variables for SET command")
				// 	}
				// 	cmd := SetCommand{
				// 		key: v.Array()[1].Bytes(),
				// 		val: v.Array()[2].Bytes(),
				// 	}

				// 	p.msgCh <- Message{
				// 		cmd: cmd,
				// 		peer: p,
				// 	}

				// 	// fmt.Printf("got SET cmd %+v\n",cmd)
				// case CommandGET:
				// 	if len(v.Array())!=2{
				// 		return fmt.Errorf("invalid no. of variables for GET command")
				// 	}
				// 	cmd := GetCommand{
				// 		key: v.Array()[1].Bytes(),
				// 	}
				// 	p.msgCh <- Message{
				// 		cmd: cmd,
				// 		peer: p,
				// 	}
				
				// case CommandHELLO:
				// 	if len(v.Array())!=2{
				// 		return fmt.Errorf("invalid no. of variables for HELLO command")
				// 	}
				// 	cmd := HelloCommand{
				// 		val: v.Array()[1].String(),
				// 	}
				// 	p.msgCh <- Message{
				// 		cmd: cmd,
				// 		peer: p,
				// 	}
				// 	// fmt.Printf("got GET cmd %+v\n",cmd)
				// default:
				// 	// panic("This command is not handled yet ...")
				// 	fmt.Println("the value string angle =>",value.String())
				// 	fmt.Printf("got unknown command => %+v\n",v.Array())
				// }
				
			// }
		}
	}

	return nil
}