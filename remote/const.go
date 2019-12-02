package remote

const (
	// VER sockts5 version  defualt  0x05  lenght 1
	VER = 0x05

	// VERIFICATIONDONT verification method
	VERIFICATIONDONT = 0x00

	// CONNECT  client request command
	CONNECT = 0X01

	// IPTYPEV4  ip v4 type
	IPTYPEV4 = 0x01

	// IPTYPEDOMAIN ip domain type
	IPTYPEDOMAIN = 0x03

	// IPTYPEV6 ip v6 type
	IPTYPEV6 = 0x04

	// RESOPNSE response sockts5 client
	RESOPNSE = 0x00
)
