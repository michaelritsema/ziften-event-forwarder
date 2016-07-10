package main

import (
	/*
		"encoding/base64"
		"fmt"

		"github.com/golang/protobuf/proto"
		"github.com/michaelritsema/ziften-event-forwarder/msg"
		"sync"

	*/
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/fatih/structs"
	"github.com/golang/protobuf/proto"
	"reflect"
	"time"
	//"time"
	_ "github.com/michaelritsema/ziften-event-forwarder/msg"
)

type AgentXMLMessage struct {
	XMLName     xml.Name `xml:"pb"`
	Payload     string   `xml:",chardata"`
	MessageType string   `xml:"type,attr"`
	Hmac        string   `xml:"HMAC,attr"`
}

// wrapper for protobuf
type AgentMessage struct {
	MessageType string
	Protobuf    interface{}
}

func WindowsTimeToUnixTime(windows_time int64) int64 {
	// number of milliseconds between Jan 1st 1601 and Jan 1st 1970
	var time_shift int64
	time_shift = 11644473600000

	if windows_time == 0 {
		return windows_time
	}

	windows_time /= 10000      // ns to ms
	windows_time -= time_shift // since 1601 to since 1970
	windows_time /= 1000
	return windows_time
}

func (m *AgentXMLMessage) toMap() map[string]interface{} {

	bytes, _ := base64.StdEncoding.DecodeString(m.Payload)
	//var messageMap map[string]interface{} = nil

	msgtype := proto.MessageType(m.MessageType).Elem()
	x := reflect.New(msgtype).Interface()
	proto.Unmarshal(bytes, x.(proto.Message))
	msgmap := structs.Map(x)

	return msgmap

	//pmsg := &protomsg.IpfixEncapsulation{}
	//proto.Unmarshal(bytes, pmsg)
}

func ParseXML(xmlmsg string) AgentXMLMessage {
	v := AgentXMLMessage{}
	xml.Unmarshal([]byte(xmlmsg), &v)
	return v
}

// sends back messages in mapChannel
func KafkaInit(mapChannel chan map[string]interface{}) {
	brokers := []string{"54.172.31.36:9092"}

	master, err := sarama.NewConsumer(brokers, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer func() {
		if err := master.Close(); err != nil {
			panic(err)
		}
	}()

	//  only support reading ipfix messages
	topic := "ZIFTEN.DATACOLLECTION_"
	// How to decide partition, is it fixed value...?
	partitions, _ := master.Partitions(topic)
	fmt.Println(partitions)
	consumers := make([]sarama.PartitionConsumer, 0)
	for _, p := range partitions {
		consumer, err := master.ConsumePartition(topic, p, sarama.OffsetOldest)
		consumers = append(consumers, consumer)
		if err != nil {
			panic(err)
		}
	}

	msgchan := make(chan *sarama.ConsumerMessage, 256)

	for x, c := range consumers {
		fmt.Printf("Init consumer partition %d\n", x)

		go func(pc sarama.PartitionConsumer) {
			msgCounter := 0

			for {
				message := <-pc.Messages()
				msgchan <- message
				msgCounter++
			}
		}(c)

		go func() {
			for msg := range msgchan {
				parsed := ParseXML(string(msg.Value[:]))
				msgmap := parsed.toMap()
				msgmap["event_type"] = parsed.MessageType
				mapChannel <- msgmap
			}
		}()
	}

}

func main() {
	c := make(chan map[string]interface{})

	KafkaInit(c)
	for {
		msg := <-c
		fmt.Println(GenerateLeef(msg))
		time.Sleep(0 * time.Second)
	}
}
