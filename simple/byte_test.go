package simple

import (
	"reflect"
	"testing"
)

func TestInt16(t *testing.T) {
	orig := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	i16 := int16fromByte(orig)
	b8 := int16toByte(i16)
	if !reflect.DeepEqual(b8, orig) {
		t.Fail()
	}
}

func TestInt32(t *testing.T) {
	orig := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	i32 := int32fromByte(orig)
	b8 := int32toByte(i32)
	if !reflect.DeepEqual(b8, orig) {
		t.Fail()
	}
}
