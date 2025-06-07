package global

const (
	AESBufferSize    = 128  // AES buffer size should be 128, 256, ...
	DESBufferSize    = 128  // DES buffer size should be 128, 256, ...
	RSAPacketSize    = 64   // RSA buffer size should less than 128(Packet)
	RSAUnpackSize    = 128  // RSA buffer size(Unpack)
	Base64BufferSize = 128  // Base64 buffer size
	ConfineFiles     = 5    // Confine go-routine concurrent files
	ConfineBuffers   = 8192 // Confine go-routine concurrent buffers
)
