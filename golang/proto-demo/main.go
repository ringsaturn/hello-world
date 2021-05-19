package main

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/valyala/bytebufferpool"
	"log"
)

func main() {
	foo := &Foo{
		A: "A",
		B: "b",
	}
	log.Println(foo)
	var ma = jsonpb.Marshaler{
		EnumsAsInts:  true,
		OrigName:     true,
		EmitDefaults: true,
	}
	buffer := bytebufferpool.Get()
	_ = ma.Marshal(buffer, foo)
	log.Println(string(buffer.Bytes()))
	bytebufferpool.Put(buffer)
}
