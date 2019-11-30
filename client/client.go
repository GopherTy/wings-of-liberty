package client

import (
	"Go_Overwall/config"
	"Go_Overwall/core"
	"Go_Overwall/encryption"
	"net"
)

// Client  cross firewall client
type Client struct {
	*core.SecoureSocket
}

// Listen listen local port
func (c *Client) Listen(didListen func(listenAdd net.Addr)) (err error) {
	listener, err := net.ListenTCP("tcp", c.ListenAddr)
	if err != nil {
		return
	}
	defer listener.Close()
	if didListen != nil {
		didListen(listener.Addr())
	}

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
		err := conn.SetLinger(0)
		if err != nil {
			conn.Close()
			sugar.Warnf("client clear data fail %v", err)
			continue
		}
		go c.handleConn(conn)
	}
}

func (c *Client) handleConn(conn *net.TCPConn) {
	defer conn.Close()
	config := config.GetConfig()
	sugar := config.Logger.Sugar()
	defer sugar.Sync()

	proxyServer, err := c.DailRemoteServer()
	defer proxyServer.Close()

	if err != nil {
		sugar.Warn("dail remote server fail %v", err)
		return
	}

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
