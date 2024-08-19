package client

import (
	// "bytes"
	"context"
	"fmt"
	"log"
	"strconv"

	// "sync"
	"testing"

	// "github.com/tidwall/resp"
)

func TestNewClient1(t *testing.T){
	c, err := New("localhost:3000")
	if err!=nil{
		log.Fatal(err)
	}
	defer c.Close()
	if err:=c.Set(context.TODO(), "foo",69);err!=nil{
		log.Fatal(err)
	}
	fmt.Println("start GET")
	val,err:=c.Get(context.TODO(), "foo")
	if err!=nil{
		log.Fatal(err)
	}
	n,_ := strconv.Atoi(val)
	fmt.Println("end GET")
	fmt.Println(n)
	
}

func TestNewClient(t *testing.T){
	c, err := New("localhost:3000")
	if err!=nil{
		log.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		
		if err:=c.Set(context.TODO(), fmt.Sprintf("foo_%d", i),fmt.Sprintf("bar_%d", i));err!=nil{
			log.Fatal(err)
		}
		fmt.Println("start GET")
		val,err:=c.Get(context.TODO(), fmt.Sprintf("foo_%d", i))
		if err!=nil{
			log.Fatal(err)
		}
		fmt.Println("end GET")
		fmt.Println(val)
	}
}