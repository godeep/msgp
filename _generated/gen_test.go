package _generated

import (
	"bytes"
	"encoding/json"
	"github.com/ugorji/go/codec"
	"reflect"
	"testing"
)

// this will work if we compile
func TestBuild(t *testing.T) {}

func BenchmarkFastEncode(b *testing.B) {
	f := &TestFast{
		Lat:  41.39082,
		Long: -41.349082,
		Alt:  201.9082,
		Data: []byte("nklasfl/kn32y0[[9uas;J1243;jsaf-0pioj124:LKas8yh34"),
	}

	var buf bytes.Buffer
	f.EncodeMsg(&buf)
	b.SetBytes(int64(len(buf.Bytes())))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		f.EncodeMsg(&buf)
	}
}

func BenchmarkFastUnmarshal(b *testing.B) {
	f := &TestFast{
		Lat:  41.39082,
		Long: -41.349082,
		Alt:  201.9082,
		Data: []byte("nklasfl/kn32y0[[9uas;J1243;jsaf-0pioj124:LKas8yh34"),
	}
	var buf bytes.Buffer
	f.EncodeMsg(&buf)
	b.SetBytes(int64(len(buf.Bytes())))
	b.ReportAllocs()
	b.ResetTimer()
	bts := buf.Bytes()
	for i := 0; i < b.N; i++ {
		f.Unmarshal(bts)
	}
}

func TestEncodeDecode(t *testing.T) {
	f := 32.00
	tt := &TestType{
		F: &f,
		Els: map[string]string{
			"thing_one": "one",
			"thing_two": "two",
		},
		Obj: struct {
			ValueA string `msg:"value_a"`
			ValueB []byte `msg:"value_b"`
		}{
			ValueA: "here's the first inner value",
			ValueB: []byte("here's the second inner value"),
		},
		Child: nil,
	}

	var buf bytes.Buffer

	_, err := tt.EncodeMsg(&buf)
	if err != nil {
		t.Fatal(err)
	}

	tnew := new(TestType)

	_, err = tnew.DecodeMsg(&buf)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(tt, tnew) {
		t.Logf("in: %v", tt)
		t.Logf("out: %v", tnew)
		t.Fatal("objects not equal")
	}
}

func BenchmarkEncodeDecode(b *testing.B) {
	f := 32.00
	tt := &TestType{
		F: &f,
		Els: map[string]string{
			"thing_one": "one",
			"thing_two": "two",
		},
		Obj: struct {
			ValueA string `msg:"value_a"`
			ValueB []byte `msg:"value_b"`
		}{
			ValueA: "here's the first inner value",
			ValueB: []byte("here's the second inner value"),
		},
		Child: nil,
	}
	var buf bytes.Buffer

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		_, err := tt.EncodeMsg(&buf)
		if err != nil {
			b.Fatal(err)
		}
		_, err = tt.DecodeMsg(&buf)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkReflectEncodeDecode(b *testing.B) {
	f := 32.00
	tt := &TestType{
		F: &f,
		Els: map[string]string{
			"thing_one": "one",
			"thing_two": "two",
		},
		Obj: struct {
			ValueA string `msg:"value_a"`
			ValueB []byte `msg:"value_b"`
		}{
			ValueA: "here's the first inner value",
			ValueB: []byte("here's the second inner value"),
		},
		Child: nil,
	}
	var buf bytes.Buffer

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		codec.NewEncoder(&buf, &codec.MsgpackHandle{}).Encode(tt)
		codec.NewDecoder(&buf, &codec.MsgpackHandle{}).Decode(tt)
	}
}

func BenchmarkJSONEncodeDecode(b *testing.B) {
	f := 32.00
	tt := &TestType{
		F: &f,
		Els: map[string]string{
			"thing_one": "one",
			"thing_two": "two",
		},
		Obj: struct {
			ValueA string `msg:"value_a"`
			ValueB []byte `msg:"value_b"`
		}{
			ValueA: "here's the first inner value",
			ValueB: []byte("here's the second inner value"),
		},
		Child: nil,
	}
	var buf bytes.Buffer

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		json.NewEncoder(&buf).Encode(tt)
		json.NewDecoder(&buf).Decode(tt)
	}
}