package main

import (
	"net"
	"strconv"
	"wings-of-liberty/config"
	"wings-of-liberty/encryption"
	"wings-of-liberty/local"
)

const (
	// DefaultListenAddr default local address
	DefaultListenAddr = ":7448"
)

func main() {
	cfg := config.GetConfig()
	sugar := cfg.Logger.Sugar()
	defer sugar.Sync()

	arr := encryption.EncrypArray{112, 44, 189, 230, 210, 98, 113, 97, 185, 211, 120, 234, 37, 106, 242, 241, 150, 1, 60, 123, 141, 19, 142, 251, 184, 195, 233, 85, 75, 226, 214, 88, 25, 202, 180, 201, 171, 116, 206, 153, 79, 23, 255, 7, 139, 51, 13, 105, 95, 2, 224, 77, 56, 114, 235, 223, 72, 119, 220, 128, 227, 43, 4, 40, 215, 252, 14, 194, 73, 193, 196, 80, 213, 11, 93, 182, 217, 87, 110, 125, 164, 5, 57, 225, 101, 82, 45, 155, 129, 199, 6, 216, 103, 231, 84, 55, 26, 168, 81, 115, 18, 192, 127, 244, 254, 0, 249, 134, 17, 169, 240, 173, 109, 27, 32, 118, 107, 130, 253, 165, 48, 136, 245, 36, 16, 159, 190, 126, 246, 146, 47, 147, 104, 54, 9, 191, 33, 59, 207, 71, 187, 89, 144, 63, 222, 99, 163, 86, 181, 143, 174, 248, 188, 10, 176, 145, 133, 250, 198, 157, 221, 212, 100, 61, 158, 94, 122, 172, 239, 132, 204, 152, 29, 46, 70, 65, 38, 151, 52, 148, 28, 41, 15, 238, 3, 117, 197, 78, 21, 74, 83, 124, 137, 208, 91, 140, 121, 138, 34, 131, 186, 232, 229, 156, 209, 30, 24, 62, 64, 167, 236, 218, 76, 35, 12, 69, 68, 178, 67, 203, 92, 31, 66, 8, 160, 243, 177, 53, 102, 90, 179, 58, 219, 49, 149, 247, 170, 22, 166, 161, 20, 135, 50, 175, 96, 237, 162, 42, 228, 205, 200, 183, 154, 108, 111, 39}

	laddrStr := cfg.Freedom.LocalAddr + ":" +
		strconv.Itoa(cfg.Freedom.LocalPort)
	raddrStr := cfg.Freedom.RemotoAddr + ":" +
		strconv.Itoa(cfg.Freedom.RemotoPort)

	laddr, err := net.ResolveTCPAddr("tcp", laddrStr)
	if err != nil {
		sugar.Fatal(err)
	}
	remote, err := net.ResolveTCPAddr("tcp", raddrStr)
	if err != nil {
		sugar.Fatal(err)
	}
	// create a client
	client := local.NewClient(&arr, laddr, remote)
	sugar.Infof("client running, local address is %s, remoto address is %s",
		laddrStr,
		raddrStr,
	)

	err = client.Listen()
	if err != nil {
		return
	}
}
