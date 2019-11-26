package encryption

import (
	"math/rand"
	"time"
)

// EncrypArray encryption array
type EncrypArray [ARRAYLEN]byte

// rand seed
func init() {
	rand.Seed(time.Now().Unix())
}

//RandEncryArray rand encryption array
func RandEncryArray() (arr *EncrypArray) {
	intArr := rand.Perm(ARRAYLEN)
	arr = &EncrypArray{}
	for i, v := range intArr {
		arr[i] = byte(v)
		if i == v {
			return RandEncryArray()
		}
	}
	return
}
