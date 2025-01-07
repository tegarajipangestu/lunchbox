package lunchbox

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func StartClient() {
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter message: ")
		message, _ := reader.ReadString('\n')
		conn.Write([]byte(message))

		response, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Server response: ", response)
	}
}
