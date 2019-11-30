package server

import (
	"Go_Overwall/config"
	"Go_Overwall/core"
	"Go_Overwall/encryption"
	"encoding/binary"
	"net"
)

// Server cross firewall server
type Server struct {
	*core.SecoureSocket
}

// Listen listening local proxy request
func (s *Server) Listen() (err error) {
	config := config.GetConfig()
	suger := config.Logger.Sugar()
	defer suger.Sync()

	listener, err := net.ListenTCP("tcp", s.ListenAddr)
	if err != nil {
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			suger.Warnf("server accept fail %v", err)
			continue
		}
		conn.SetLinger(0)
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn *net.TCPConn) (err error) {
	defer conn.Close()
	buf := make([]byte, encryption.ARRAYLEN)

	// recive socket protocol
	/**
	   The localConn connects to the dstServer, and sends a ver
	   identifier/method selection message:
		          +----+----------+----------+
		          |VER | NMETHODS | METHODS  |
		          +----+----------+----------+
		          | 1  |    1     | 1 to 255 |
		          +----+----------+----------+
	   The VER field is set to X'05' for this ver of the protocol.  The
	   NMETHODS field contains the number of method identifier octets that
	   appear in the METHODS field.
	*/
	_, err = s.DecryptData(conn, buf)
	if err != nil || buf[0] != VER {
		return
	}
	/**
	   The dstServer selects from one of the methods given in METHODS, and
	   sends a METHOD selection message:

		          +----+--------+
		          |VER | METHOD |
		          +----+--------+
		          | 1  |   1    |
		          +----+--------+
	*/
	// don't verification
	_, err = s.EncryptData(conn, []byte{VER, VERIFICATIONDONT})
	if err != nil {
		return
	}
	// client and server connected, receive encrypted data which is contanted real remote address
	/**
	  +----+-----+-------+------+----------+----------+
	  |VER | CMD |  RSV  | ATYP | DST.ADDR | DST.PORT |
	  +----+-----+-------+------+----------+----------+
	  | 1  |  1  | X'00' |  1   | Variable |    2     |
	  +----+-----+-------+------+----------+----------+
	*/
	// get remote address
	n, err := s.DecryptData(conn, buf)
	if err != nil || n < 7 {
		return
	}

	// get client request command
	if buf[1] != CONNECT {
		// only surpport connect
		return
	}

	// the type of  remote ip, lenght  1
	var remoteIP []byte
	var domain *net.IPAddr
	switch buf[3] {
	case IPTYPEV4:
		remoteIP = buf[4 : 4+net.IPv4len]
	case IPTYPEDOMAIN:
		domain, err = net.ResolveIPAddr("ip", string(buf[5:n-2]))
		if err != nil {
			return
		}
		remoteIP = domain.IP
	case IPTYPEV6:
		remoteIP = buf[4 : 4+net.IPv4len]
	default:
		return
	}
	remotePort := buf[n-2:]
	remoteAddr := &net.TCPAddr{
		IP:   remoteIP,
		Port: int(binary.BigEndian.Uint16(remotePort)),
	}

	// real remote address
	server, err := net.DialTCP("tcp", nil, remoteAddr)
	if err != nil {
		return
	}
	defer server.Close()
	server.SetLinger(0)

	// response client
	/**
	  +----+-----+-------+------+----------+----------+
	  |VER | REP |  RSV  | ATYP | BND.ADDR | BND.PORT |
	  +----+-----+-------+------+----------+----------+
	  | 1  |  1  | X'00' |  1   | Variable |    2     |
	  +----+-----+-------+------+----------+----------+
	*/
	_, err = s.EncryptData(conn, []byte{VER, RESOPNSE, RESOPNSE, CONNECT, RESOPNSE, RESOPNSE, RESOPNSE, RESOPNSE, RESOPNSE, RESOPNSE})
	if err != nil {
		return
	}
	go func() {
		err = s.DecryptCopy(server, conn)
		if err != nil {
			conn.Close()
			server.Close()
		}
	}()
	s.EncryptCopy(conn, server)
	return
}

// NewServer create a server
func NewServer(arr *encryption.EncrypArray, listenAddr *net.TCPAddr) (server *Server) {
	server = &Server{
		SecoureSocket: &core.SecoureSocket{
			Cipher:     encryption.NewCipher(arr),
			ListenAddr: listenAddr,
		},
	}
	return
}
