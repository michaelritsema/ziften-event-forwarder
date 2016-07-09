package main

import (
	proto "github.com/golang/protobuf/proto"
	msg "github.com/michaelritsema/ziften-event-forwarder/msg"
	"testing"
)

func TestProto(t *testing.T) {

	fh := msg.FrameHeader{}
	fh.Length = proto.Int32(1)
	fh.Version = proto.Int32(2)
	fh.Type = proto.String("x")

}
