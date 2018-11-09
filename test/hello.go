package main

import (
	"bufio"
	"fmt"
	//"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	VERSION = iota
	APP
	INVOICE
)

func main() {
	service := ":10000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkErr(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkErr(err)
	fmt.Printf("start server..")
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	var version int
	f, err := os.Open("version")
	defer f.Close()
	if err == nil {
		buff := bufio.NewReader(f)
		line, err := buff.ReadString('\n')
		if err == nil {
			str := strings.Split(line, "=")[1]
			str = strings.Replace(str, "\n", "", -1)
			version, err = strconv.Atoi(str)
			fmt.Println("version", version)
		}
	}
	/*
		var buf [100]byte
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		rAddr := conn.RemoteAddr()
		fmt.Println("Receive from client", rAddr.String(), string(buf[0:n]))
	*/
	vbuf := append([]byte{}, byte(VERSION), byte(version))
	fmt.Println("vBuf", vbuf)
	_, err2 := conn.Write(vbuf)
	if err2 != nil {
		fmt.Println("conn write error", err2)
		return
	}
	var buf [1]byte
	n, err := conn.Read(buf[0:1])
	fmt.Println("read ", err, n)
	if err != nil {
		fmt.Println("version request read Error")
	}
	if buf[0] == 1 {
		fmt.Println("version request ", buf[0])
	}
}
