package main

import (
	"testing"
)

func TestGenerateHeader(t *testing.T) {
	t.Logf(generateHeader("ProcessInventory"))
}

func TestGenerateEventAttr(t *testing.T) {
	msgmap := make(map[string]interface{})
	msgmap["TimeStamp"] = 131122061909468384
	msgmap["ImageFilepath"] = `ImageFilepath=c:\windows\system32\gpsvc.dll`
	msgmap["BootAnalysisGUID"] = "{B7FA08ED-B055-427C-8811-9EDB674E2164}"
	t.Logf(generateEventAttr(msgmap))
}
