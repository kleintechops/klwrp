package main
	
import (
"os"

"time"
"io"
"fmt"
"net"
)


func flush_input_buffer(conn net.Conn) string {

fmt.Println("Flushing input buffer...")

var outString string

buffer := make([]byte, 1024)
for {
		conn.SetReadDeadline(time.Now().Add(1 * time.Second)) 
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Connection closed by remote host.")
					panic("Error flushing buffer")
			}else if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			            fmt.Println("No new data available.")
				break
		        } else {
				fmt.Println("Error reading data:", err)
					panic("Error flushing buffer")
			}
		}

		fmt.Printf("[RX-Line:]>> %s \n", buffer[:n])

		outString += string(buffer[:n])

}

		fmt.Println("[END MSG]")

		return outString

//		return string(buffer[:n])


}


func main() {

port := "93" // livewire routing (lwrp) port

// Args needed: <Node IP> <DST #> <LW Src # or MCAST IP> <FRIENDLY NAME OF DST>

if len(os.Args) != 5 {
fmt.Println("Need 4 arguments to run:\n" + os.Args[0] + " <NODE IP> <DST #> <LW SRC #> <NAME>")
os.Exit(3)
}


host := os.Args[1] // ip of node
dst := os.Args[2] // destination channel of node
lwsrc := os.Args[3] // source to route to node dst
name := os.Args[4] // friendly name for dst channel

var configString string
configString = "DST " + dst + " NAME:\"" + name + "\" ADDR:" + lwsrc + "\n"

conn, err := net.Dial("tcp", host+":"+port)

if err != nil {
	panic("Error starting connection!")
}

fmt.Println("Connection started")
fmt.Println(configString)

loginString := "LOGIN \n"
conn.Write([]byte(loginString))
fmt.Println("Sent LOGIN")

fmt.Println(flush_input_buffer(conn))


conn.Write([]byte(configString))
fmt.Println("Sent configuration.")

fmt.Println(flush_input_buffer(conn))


os.Exit(0)

}
