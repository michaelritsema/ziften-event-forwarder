package main

import (
	"encoding/base64"
	"github.com/golang/protobuf/proto"
	_ "github.com/michaelritsema/ziften-event-forwarder/msg"
	"reflect"
	"testing"
)

var testxml string = `<pb type="Services" HMAC="RiTl7K5Cr+4oH0SHzWdumLJBua4=">CODJ0pKh2vXoARImezJCQUI5QjUwLTNBRjUtMTFFNi05QkVCLTAwNTA1NkUzQkFFQX0aBWdwc3ZjIhNHcm91cCBQb2xpY3kgQ2xpZW50KipDOlxXaW5kb3dzXHN5c3RlbTMyXHN2Y2hvc3QuZXhlIC1rIG5ldHN2Y3MyHWM6XHdpbmRvd3Ncc3lzdGVtMzJcZ3BzdmMuZGxsOCBAAkogYjY4NWE5ZDA1YjVlY2MzYmViNjY0YzdjNDYwYmM4NWNSIDdjZDY5ODI2ZjZkOTFmZWVjMjU3NzQyMjIyNjI2Y2MzWiZ7NDg2MTJGN0ItM0NFRS00RjFGLUJGRjUtMUZCOTY3ODRDQzhDfWIme0I3RkEwOEVELUIwNTUtNDI3Qy04ODExLTlFREI2NzRFMjE2NH0=</pb>`

func TestParseXML(t *testing.T) {
	unmarshalled := ParseXML(testxml)
	t.Logf("XMLName: %s , messageType :%s, HMAC: %s ,payload: \n %s", unmarshalled.XMLName, unmarshalled.MessageType, unmarshalled.Hmac, unmarshalled.Payload)
}

func TestToMap(t *testing.T) {
	unmarshalled := ParseXML(testxml)
	t.Logf("Map:\n")
	for k, v := range unmarshalled.toMap() {
		if k != "XXX_unrecognized" {
			t.Logf("%v=%v", k, reflect.Indirect(reflect.ValueOf(v)))
		}
	}

}

func TestCreateProtobufFromReflection(t *testing.T) {
	xmlmsg := ParseXML(testxml)
	bytes, _ := base64.StdEncoding.DecodeString(xmlmsg.Payload)
	msgtype := proto.MessageType(xmlmsg.MessageType).Elem()
	x := reflect.New(msgtype).Interface()
	proto.Unmarshal(bytes, x.(proto.Message))
	t.Log(x)
}
