package blivedm

import (
	"bytes"
	"encoding/binary"
)

type WsHeader struct {
	PacketLength    uint32
	HeaderLength    uint16
	ProtocolVersion uint16
	Operation       uint32
	Sequence        uint32
}

func ResolveWSPacket(data []byte) (WsHeader, []byte, bool) {
	if len(data) < 16 {
		return WsHeader{}, []byte{}, false
	}
	header := WsHeader{
		PacketLength:    binary.BigEndian.Uint32(data[0:4]),
		HeaderLength:    binary.BigEndian.Uint16(data[4:6]),
		ProtocolVersion: binary.BigEndian.Uint16(data[6:8]),
		Operation:       binary.BigEndian.Uint32(data[8:12]),
		Sequence:        binary.BigEndian.Uint32(data[12:16]),
	}
	return header, data[header.HeaderLength:header.PacketLength], true
}

func MakeWSPacket(operation int, data []byte) []byte {
	headerBytes := new(bytes.Buffer)
	header := []interface{}{
		uint32(len(data) + 16),
		uint16(16),
		uint16(1),
		uint32(operation),
		uint32(1),
	}
	for _, v := range header {
		err := binary.Write(headerBytes, binary.BigEndian, v)
		if err != nil {
		}
	}
	return append(headerBytes.Bytes(), data...)
}
