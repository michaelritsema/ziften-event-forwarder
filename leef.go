package main

import (
	"bytes"
	"fmt"
)

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

	for k, v := range msgmap {
		buffer.WriteString(k)
		buffer.WriteString(equals)
		buffer.WriteString(fmt.Sprintf("%v", v))
		buffer.WriteString(tab)
	}

	return buffer.String()
}
