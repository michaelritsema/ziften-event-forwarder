package main

import (
	"bytes"
	"fmt"
	"time"
)

// data adjusting that should happen across all messages
func preprocessMap(msgmap map[string]interface{}) {
	msgmap["TimeStamp"] = time.Unix(WindowsTimeToUnixTime(*msgmap["TimeStamp"].(*int64)), 0).String()
}

var aliasMap = map[string]string{
	"TimeStamp": "devTime",
}

func renameFields(msgmap map[string]interface{}) {
	// Most inreresting QRadar fields that can help with IOCs
	// src (source ip)
	// dst (destination ip)
	// proto (protocol)
	// srcPort
	// dstPort

	// Other QRadar normalized fields
	// devTime
	for k, v := range aliasMap {
		_, ok := msgmap[k]
		if ok {
			msgmap[v] = msgmap[k]
			delete(msgmap, k)
		}
	}
}

// adjust map keys and values for LEEF (QRadar)
func formatMap(msgmap map[string]interface{}) {
	preprocessMap(msgmap)
	renameFields(msgmap)
}

func generateHeader(eventType string) string {
	leefVersion := "1.0"
	productVendorName := "Ziften"
	productName := "Ziften EDR"
	ziftenVersion := "5.0.1"

	return fmt.Sprintf("LEEF:%s|%s|%s|%s|%s|", leefVersion, productVendorName, productName, ziftenVersion, eventType)
}

func generateEventAttr(msgmap map[string]interface{}) string {
	var buffer bytes.Buffer
	var equals = "="
	var tab = "\t"
	formatMap(msgmap)
	for k, v := range msgmap {
		buffer.WriteString(k)
		buffer.WriteString(equals)
		buffer.WriteString(fmt.Sprintf("%v", v))
		buffer.WriteString(tab)
	}

	return buffer.String()
}
