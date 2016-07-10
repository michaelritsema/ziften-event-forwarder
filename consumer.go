package main

import (
	/*
		"encoding/base64"
		"fmt"
		"github.com/Shopify/sarama"
		"github.com/golang/protobuf/proto"
		"github.com/michaelritsema/ziften-event-forwarder/msg"
		"sync"

	*/
	"encoding/base64"
	"encoding/xml"
	"github.com/fatih/structs"
	"github.com/golang/protobuf/proto"
	"reflect"
	"time"
	//"github.com/michaelritsema/ziften-event-forwarder/msg"
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

// data adjusting that should happen across all messages
func preprocessMap(msgmap map[string]interface{}) {
	msgmap["TimeStamp"] = time.Unix(WindowsTimeToUnixTime(*msgmap["TimeStamp"].(*int64)), 0).String()
}

func (m *AgentXMLMessage) toMap() map[string]interface{} {

	bytes, _ := base64.StdEncoding.DecodeString(m.Payload)
	//var messageMap map[string]interface{} = nil

	msgtype := proto.MessageType(m.MessageType).Elem()
	x := reflect.New(msgtype).Interface()
	proto.Unmarshal(bytes, x.(proto.Message))
	msgmap := structs.Map(x)
	preprocessMap(msgmap)
	return msgmap

	//pmsg := &protomsg.IpfixEncapsulation{}
	//proto.Unmarshal(bytes, pmsg)
}

func ParseXML(xmlmsg string) AgentXMLMessage {
	v := AgentXMLMessage{}
	xml.Unmarshal([]byte(xmlmsg), &v)
	return v
}

func main() {

}

/*
	takes in <pb> </pb>
	returns array of ipfix messages as [][]byte
*/

/*
func stripXML(msg []byte) [][]byte {
	payload := msg[66 : len(msg)-5]
	bytes, _ := base64.StdEncoding.DecodeString(string(payload))
	pmsg := &protomsg.IpfixEncapsulation{}
	proto.Unmarshal(bytes, pmsg)

	return pmsg.GetIpfixMessage()
}

func Kafka(protoMsgChan chan *protomsg.ZFlow, connection string) {
	converter := iespec.InitIpfix()

	//config := sarama.NewConfig()

	//config.Consumer.Return.Errors = true

	// Specify brokers address. This is default one
	//brokers := []string{"ec2-107-21-70-96.compute-1.amazonaws.com:9092"}
	brokers := []string{connection}
	fmt.Println(brokers)
	//ec2-54-210-70-189.compute-1.amazonaws.com:9092
	// Create new coeeensumer
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
	topic := "ZIFTEN.IPFIX"
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

				//fmt.Println(msgCounter)
			}
		}(c)

		go func() {
			for msg := range msgchan {
				for _, ipfixBytes := range stripXML(msg.Value) {

					//fmt.Printf("Value:\t%v\n", ipfixBytes)

					for _, zflowMsg := range converter.Convert(ipfixBytes) {
						//fmt.Println(zflowMsg)
						protoMsgChan <- zflowMsg
						//fmt.Printf("Partition:\t%d\n", msg.Partition)
						//fmt.Printf("Offset:\t%d\n", msg.Offset)
						//fmt.Printf("Key:\t%s\n", string(msg.Key))

						//fmt.Println()

					}
				}

			}
		}()
	}
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}

/*
func main() {
	c := make(chan *protomsg.ZFlow)
	go kafka(c)
	for {
		x := <-c
		fmt.Println(x)
	}
}

*/
