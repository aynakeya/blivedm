package blivedm

const (
	// OpHandshake handshake
	OpHandshake = 0
	// OpHandshakeReply handshake reply
	OpHandshakeReply = 1

	// OpHeartbeat heartbeat
	OpHeartbeat = 2
	// OpHeartbeatReply heartbeat reply
	OpHeartbeatReply = 3

	// OpSendMsg send message.
	OpSendMsg = 4
	// OpSendMsgReply  send message reply
	OpSendMsgReply = 5

	// OpDisconnectReply disconnect reply
	OpDisconnectReply = 6

	// OpAuth auth connnect
	OpAuth = 7
	// OpAuthReply auth connect reply
	OpAuthReply = 8

	// OpRaw  raw message
	OpRaw = 9

	// OpProtoReady proto ready
	OpProtoReady = 10
	// OpProtoFinish proto finish
	OpProtoFinish = 11

	// OpChangeRoom change room
	OpChangeRoom = 12
	// OpChangeRoomReply change room reply
	OpChangeRoomReply = 13

	// OpRegister register operation
	OpRegister = 14
	// OpRegisterReply register operation
	OpRegisterReply = 15

	// OpUnregister unregister operation
	OpUnregister = 16
	// OpUnregisterReply unregister operation reply
	OpUnregisterReply = 17

	// MinBusinessOp min business operation
	MinBusinessOp = 1000
	// MaxBusinessOp max business operation
	MaxBusinessOp = 10000
)
