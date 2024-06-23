package frame


type FrameType int

const (
	ACK FrameType = iota
	CONNECTION_CLOSE
	CRYPTO
	DATA_BLOCKED
	HANDSHAKE_DONE
	MAX_DATA
	MAX_STREAM_DATA
	MAX_STREAMS
	NEW_CONNECTION_ID
	NEW_TOKEN
	PADDING
	PATH_CHALLENGE
	PATH_RESPONSE
	PING
	RESET_STREAM
	RETIRE_CONNECTION_ID
	STOP_SENDING
	STREAM
	STREAM_DATA_BLOCKED
	STREAMS_BLOCKED
)


// 1-RTT Packet Format
// 1-RTT Packet {
// 	Header Form (1) = 0,
// 	Fixed Bit (1) = 1,
// 		- The next bit (0x40) of byte 0 is set to 1. Packets containing a zero value for this bit are not valid packets in this version and MUST be discarded. A value of 1 for this bit allows QUIC to coexist with other protocols; see [RFC7983].
// 	Spin Bit (1),
// 		- Enables passive latency monitoring for observation points on the network path throughout the duration of a connection
// 		- spin bit is optional
//		- spin bits are controled either globally or per endpoint
// 	Reserved Bits (2),
// 	Key Phase (1),
// 	Packet Number Length (2),
// 	Destination Connection ID (0..160),
// 	Packet Number (8..32),
// 	Packet Payload (8..),
// }
type RTT1Packet struct {}
