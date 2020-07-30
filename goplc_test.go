package goplc

import (
	"bytes"
	"fmt"
	"github.com/MiguelValentine/goplc/enip/cip/dataTypes"
	"github.com/MiguelValentine/goplc/enip/cip/epath/segments/data"
	"github.com/MiguelValentine/goplc/enip/cip/epath/segments/logical"
	"github.com/MiguelValentine/goplc/enip/cip/epath/segments/port"
	"testing"
)

// Valid Range 0xc1 - 0xdf
func TestIsValidTypeCode(t *testing.T) {
	if dataTypes.IsValidTypeCode(0x11) {
		t.Error("Invalid Code Check Error")
	}
	if !dataTypes.IsValidTypeCode(0xc1) {
		t.Error("Valid Code Check Error")
	}
}

func TestGetTypeCodeString(t *testing.T) {
	if dataTypes.GetTypeCodeString(0xc1) != "BOOL" {
		t.Error()
	}
}

func TestLogicalBuild(t *testing.T) {
	bf1 := logical.Build(logical.ClassID, 15, true)
	bf2 := logical.Build(logical.ClassID, 1755, true)
	bf3 := logical.Build(logical.ClassID, 71755, true)

	if fmt.Sprintf("%x", bf1.Bytes()) != "200f" {
		t.Error()
	}
	if fmt.Sprintf("%x", bf2.Bytes()) != "2100db06" {
		t.Error()
	}
	if fmt.Sprintf("%x", bf3.Bytes()) != "22004b180100" {
		t.Error()
	}
}

func TestPortBuild(t *testing.T) {
	e1, bf1 := port.Build(1, "aa")
	if e1 != nil {
		t.Error(e1)
	}

	if fmt.Sprintf("%x", bf1.Bytes()) != "11026161" {
		t.Error()
	}

	e2, bf2 := port.Build(7, uint8(8))
	if e2 != nil {
		t.Error(e2)
	}

	if fmt.Sprintf("%x", bf2.Bytes()) != "0708" {
		t.Error()
	}

	e3, bf3 := port.Build(16, "aax")
	if e3 != nil {
		t.Error(e3)
	}

	if fmt.Sprintf("%x", bf3.Bytes()) != "1f03100061617800" {
		t.Error()
	}

	e4, bf4 := port.Build(25, uint8(9))
	if e4 != nil {
		t.Error(e4)
	}

	if fmt.Sprintf("%x", bf4.Bytes()) != "0f190009" {
		t.Error()
	}
}

func TestDataBuild(t *testing.T) {
	bf1 := data.Build(uint32(15), true)
	if bf1 == nil || fmt.Sprintf("%x", bf1.Bytes()) != "280f" {
		t.Error()
	}

	bf2 := data.Build(uint32(257), true)
	if bf2 == nil || fmt.Sprintf("%x", bf2.Bytes()) != "29000101" {
		t.Error()
	}

	bf3 := data.Build(uint32(65534), true)
	if bf3 == nil || fmt.Sprintf("%x", bf3.Bytes()) != "2900feff" {
		t.Error()
	}

	bf4 := data.Build(uint32(65540), true)
	if bf4 == nil || fmt.Sprintf("%x", bf4.Bytes()) != "2a0004000100" {
		t.Error()
	}

	bf5 := data.Build(bytes.NewBufferString("hello"), true)
	if bf5 == nil || fmt.Sprintf("%x", bf5.Bytes()) != "910568656c6c6f00" {
		t.Error()
	}

	bf6 := data.Build(bytes.NewBufferString("hello"), false)
	if bf6 == nil || fmt.Sprintf("%x", bf6.Bytes()) != "800368656c6c6f00" {
		t.Error()
	}
}

func TestMessageBuild(t *testing.T) {

}