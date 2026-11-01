package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func startServer(port string) {

	//	CONCATENATING ":" INTO THE PORT AS THE LISTENER WANTS THE PORT IN THE FORMAT :PORT
	port = ":" + port

	//	STARTING THE LISTENER
	listener, err := net.Listen("tcp", port)

	//	CHECKING FOR ERROR, IF ANY PRINT IT OUT AND EXIT
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//	DEFERING THE LISTENER TO BE CLOSED AT THE END
	defer listener.Close()

	//	STARTING THE SOCKET WHICH ACCEPTS THE INCOMING CONNECTIONS
	socket, err := listener.Accept()

	//	CHECKING FOR ERROR, IF ANY PRINT IT OUT AND EXIT
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//	DEFERING THE SOCKET TO BE CLOSED AT THE END
	defer socket.Close()

	//	CALLING THE CONNECTION HANDLER THAT SHARES THE DATA
	connHandler(socket)
}

func startClient(ip, port string) {

	// CONCATENATING THE IP AND PORT IN THE FORMAT IP:PORT SINCE THE DIAL WANTS THIS FORMAT
	ip_port := ip + ":" + port

	//	DIAL THE CONNECTION TO SPECIFIC IP AND THE PORT
	dialer, err := net.Dial("tcp", ip_port)

	//	CHECK FOR ERROR, IF ANY PRINT IT OUT AND EXIT
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//	DEFERING THE DIALER TO BE CLOSED AT THE END
	defer dialer.Close()

	//	CALLING THE CONNECTION HANDLER THAT SHARES THE DATA
	connHandler(dialer)
}

func connHandler(conn net.Conn) {

	fmt.Println("Ready to chat!")

	//	TAKING THE INPUT FROM ONE SIDE AND THROWING TO THE OTHER
	go io.Copy(conn, os.Stdin)
	io.Copy(os.Stdout, conn)
}

func main() {

	//	CREATING A CHANNEL THAT LISTENS FOR CTRL+C TO EXIT THE PROGRAM
	exitChannel := make(chan os.Signal, 1)
	signal.Notify(exitChannel, os.Interrupt, syscall.SIGTERM)

	//	STARTING A GO ROUTINE THAT WILL EXIT OUT WHEN CTRL+C
	go func() {
		<-exitChannel
		fmt.Println("\n[Ctrl+C] Exiting the session")
		os.Exit(0)
	}()

	//	CHEKS IF THE NO. OF ARGUMENTS PROVIDED MATCH THE NO. OF ARGUMENTS REQUIRED BY THE PROGRAM
	if len(os.Args) < 3 {
		fmt.Println("Usage:\n./gonc -l <PORT>\n./gonc <IP> <PORT>")
		os.Exit(1)
	}

	//	CHECKS IF THE FIRST ARGUMENT PROVIDED MATCHES WITH THE LISTENER'S BLUEPRINT
	if os.Args[1] == "-l" {

		//	EXECUTE THE SERVER FUNCTION WHICH HANDLES THE SERVER LOGIC
		startServer(os.Args[2])
	} else {

		// EXECUTES THE CLIENT FUNCTION WHICH WILL HANDLE THE CLIENT LOGIC, AND ALSO HANDLE THE SYNTAX ERROR
		startClient(os.Args[1], os.Args[2])
	}

}
