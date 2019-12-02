package core

import (
	"io"
	"net"
	"wings-of-liberty/encryption"
)

// SecoureSocket encrypt structure  which is socket data
type SecoureSocket struct {
	Cipher     *encryption.Cipher
	ListenAddr *net.TCPAddr
	RemoteAddr *net.TCPAddr
}

// DecryptData read  encryption data from input stream
func (s *SecoureSocket) DecryptData(conn *net.TCPConn, buf []byte) (n int, err error) {
	// read encryption data from input stream into buffer slice
	n, err = conn.Read(buf)
	if err != nil {
		return
	}
	// decryp data into buffer slice
	s.Cipher.Decrypt(buf[:n])
	return
}

// EncryptData .
func (s *SecoureSocket) EncryptData(conn *net.TCPConn, buf []byte) (n int, err error) {
	// encrypt data into buffer slice
	s.Cipher.Decrypt(buf)
	// write encryption data into output stream
	n, err = conn.Write(buf)
	if err != nil {
		return
	}
	return
}

// EncryptCopy .
func (s *SecoureSocket) EncryptCopy(dst, src *net.TCPConn) (err error) {
	var r, w int
	buf := make([]byte, encryption.ARRAYLEN)
	for {
		r, err = src.Read(buf)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return
		}
		if r <= 0 {
			return
		}
		w, err = s.EncryptData(dst, buf[:r])
		if err != nil {
			return
		}
		if w != r {
			return io.ErrShortWrite
		}
	}
}

// DecryptCopy .
func (s *SecoureSocket) DecryptCopy(dst, src *net.TCPConn) (err error) {
	var r, w int
	buf := make([]byte, encryption.ARRAYLEN)
	for {
		r, err = s.DecryptData(src, buf)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return
		}
		if r <= 0 {
			return
		}
		w, err = dst.Write(buf[:r])
		if err != nil {
			return
		}
		if r != w {
			return io.ErrShortWrite
		}
	}
}

// DailRemoteServer .
func (s *SecoureSocket) DailRemoteServer() (remoteConn *net.TCPConn, err error) {
	remoteConn, err = net.DialTCP("tcp", nil, s.RemoteAddr)
	if err != nil {
		return
	}
	return
}
