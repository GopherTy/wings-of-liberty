package local

import (
	"net"
	"os"
	"wings-of-liberty/config"
	"wings-of-liberty/core"
	"wings-of-liberty/encryption"
)

// Client  cross firewall client
type Client struct {
	*core.SecoureSocket
}

// Listen listen local port
func (c *Client) Listen() (err error) {
	listener, err := net.ListenTCP("tcp", c.ListenAddr)
	if err != nil {
		return
	}
	defer listener.Close()

	var conn *net.TCPConn
	config := config.GetConfig()
	sugar := config.Logger.Sugar()
	defer sugar.Sync()

	for {
		conn, err = listener.AcceptTCP()
		if err != nil {
			sugar.Warnf("client accept fail %v", err)
			continue
		}
		if conn == nil {
			return
		}
		conn.SetLinger(0)
		go c.handleConn(conn)
	}
}

func (c *Client) handleConn(conn *net.TCPConn) {
	config := config.GetConfig()
	sugar := config.Logger.Sugar()

	defer sugar.Sync()
	defer conn.Close()
	proxyServer, err := c.DailRemoteServer()
	if err != nil {
		sugar.Warn(err)
		os.Exit(1)
	}

	defer proxyServer.Close()

	go func() {
		err = c.DecryptCopy(conn, proxyServer)
		if err != nil {
			conn.Close()
			proxyServer.Close()
		}
	}()

	c.EncryptCopy(proxyServer, conn)
}

// NewClient  create a client
func NewClient(arr *encryption.EncrypArray, listenAddr, remoteAddr *net.TCPAddr) (client *Client) {
	client = &Client{
		SecoureSocket: &core.SecoureSocket{
			Cipher:     encryption.NewCipher(arr),
			ListenAddr: listenAddr,
			RemoteAddr: remoteAddr,
		},
	}
	return
}
