package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"
	"strings"
)

var (
	hostname = "localhost"
	port     = "8090"
	address  = fmt.Sprintf("%s:%s", hostname, port)
)

func main() {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("listen on %s", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		log.Printf("connected with %s", conn.RemoteAddr())

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			io.WriteString(conn, fmt.Errorf("read from conn failed: %w", err).Error())
			continue
		}
		msg = strings.TrimSuffix(msg, "\n")
		if msg == "close" {
			conn.Close()
			return
		}

		command, args := parseMsg(msg)
		cmd := exec.Command(command, args...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			io.WriteString(conn, fmt.Errorf("exec %s failed: %w", msg, err).Error())
		}

		io.WriteString(conn, string(output))
	}
}

func parseMsg(msg string) (string, []string) {
	vals := strings.Split(msg, " ")
	cmd := vals[0]
	args := vals[1:]
	return cmd, args
}
