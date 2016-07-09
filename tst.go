package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/michaelritsema/ziften-event-forwarder/msg"
)

func main() {

	fh := msg.FrameHeader{}
	fh.Length = proto.Int32(1)
	fh.Version = proto.Int32(2)
	fh.Type = proto.String("x")
	fmt.Println(fh)

}
