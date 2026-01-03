package server

import (
	"io"
	"log"
	"net"
	"strconv"

	"github.com/oyesaurabh/gedis/config"
)

func readCommand(c net.Conn) (string, error) {
	var buf []byte = make([]byte, 512)
	n, err := c.Read(buf[:])
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}
func respond(cmd string, c net.Conn) error {
	if _, err := c.Write([]byte(cmd)); err != nil {
		return err
	}
	return nil
}
func RunSyncTCPServer() {
	log.Println("starting a synchronous TCP server on", config.Host, config.Port)

	var con_clients int = 0

	// listening to the configured host:port
	lsnr, err := net.Listen("tcp", config.Host+":"+strconv.Itoa(config.Port))
	if err != nil {
		log.Fatalln("error starting the server:", err)
		return
	}

	for {
		//blocking call - waits for a new connection
		conn, err := lsnr.Accept()
		if err != nil {
			panic(err)
		}
		con_clients++
		log.Println("client address:", conn.RemoteAddr().String(), "total clients:", con_clients)

		// talking to ONE client
		for {
			cmd, err := readCommand(conn)
			if err != nil {
				if err == io.EOF {
					log.Println("client disconnected:", conn.RemoteAddr())
				} else {
					log.Println("read error:", err)
				}

				conn.Close()
				con_clients--
				log.Println("total clients:", con_clients)
				break
			}

			log.Println("command:", cmd)

			if err = respond(cmd, conn); err != nil {
				log.Println("write error:", err)
				conn.Close()
				con_clients--
				break
			}
		}

	}
}
