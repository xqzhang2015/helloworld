package main

type MyStruct struct {
	Data []byte
}

func (ms *MyStruct) GetData() []byte {
	return ms.Data
}

var Codec = "myCodec"

// MyCodec defines the interface gRPC uses to encode and decode messages.  Note
// that implementations of this interface must be thread safe; a Codec's
// methods can be called from concurrent goroutines.
type MyCodec struct{}

// Marshal returns the wire format of v.
func (MyCodec) Marshal(v interface{}) ([]byte, error) {
	return v.([]byte), nil
}

// Unmarshal parses the wire format into v.
func (MyCodec) Unmarshal(data []byte, v interface{}) error {
	ms := v.(*MyStruct)
	ms.Data = data
	return nil
}

// String  old gRPC Codec interface func
func (MyCodec) String() string {
	return Codec
}

// Name returns the name of the Codec implementation. The returned string
// will be used as part of content type in transmission.  The result must be
// static; the result cannot change between calls.
//
// add Name() for ForceCodec interface
func (MyCodec) Name() string {
	return Codec
}
