package encryption

// Cipher encryption cipher
type Cipher struct {
	encryptArray *EncrypArray
	decryptArray *EncrypArray
}

// Encrypt encrypt data by encryption array
func (cipher *Cipher) Encrypt(arr []byte) {
	for i, v := range arr {
		arr[i] = cipher.encryptArray[v]
	}
}

// Decrypt decrypt data by encryption array
func (cipher *Cipher) Decrypt(arr []byte) {
	for i, v := range arr {
		arr[i] = cipher.decryptArray[v]
	}
}

// NewCipher create a cipher structure for encrypt and decrypt data
func NewCipher(arr *EncrypArray) (c *Cipher) {
	decryptArray := &EncrypArray{}
	for i, v := range arr {
		decryptArray[v] = byte(i)
	}
	c = new(Cipher)
	c.encryptArray = arr
	c.decryptArray = decryptArray
	return
}
