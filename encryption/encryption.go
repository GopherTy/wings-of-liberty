package encryption

// Cipher encryption cipher
type Cipher struct {
	encodeEncrypArray *EncrypArray
	decodeEncrypArray *EncrypArray
}

// Encode encode data by encryption array
func (cipher *Cipher) Encode(arr []byte) {
	for i, v := range arr {
		arr[i] = cipher.encodeEncrypArray[v]
	}
}

// Decode decode data by encryption array
func (cipher *Cipher) Decode(arr []byte) {
	for i, v := range arr {
		arr[i] = cipher.decodeEncrypArray[v]
	}
}

// NewCipher create a cipher for encode and decode data
func NewCipher(arr *EncrypArray) (c *Cipher) {
	decodeEncrypArray := &EncrypArray{}
	for i, v := range arr {
		arr[i] = v
		decodeEncrypArray[v] = byte(i)
	}
	c = new(Cipher)
	c.encodeEncrypArray = arr
	c.decodeEncrypArray = decodeEncrypArray
	return
}
