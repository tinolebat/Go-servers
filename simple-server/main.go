package main

import (
	"fmt"
	"log"
	"net"
)

type Message struct {
	from    string
	payload []byte
}

type Server struct {
	listenAddr string
	ln         net.Listener
	qch        chan struct{}
	msch       chan Message
}

func NewServer(address string) *Server {
	return &Server{
		listenAddr: address,
		qch:        make(chan struct{}),
		msch:       make(chan Message, 10),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		fmt.Println("Fail to Start")
		return err
	}
	defer ln.Close()
	s.ln = ln

	fmt.Println("Server Started, listening to port", s.listenAddr)

	go s.accept()

	<-s.qch
	close(s.msch)

	return nil

}

func (s *Server) accept() {

	for {
		conn, err := s.ln.Accept()
		if err != nil {
			log.Println("Fail to Accept Conn.", err)
		}

		fmt.Printf("New Connection to server from (%s)", conn.RemoteAddr())

		go s.read(conn)

	}
}

func (s *Server) read(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read Buffer error: ", err)
			continue
		}

		s.msch <- Message{
			from:    conn.RemoteAddr().String(),
			payload: buf[:n],
		}
		conn.Write([]byte("Message received"))

	}

}

func main() {

	server := NewServer(":3000")

	go func() {
		for msg := range server.msch {
			fmt.Printf("Received message from (%s) : %s \n", msg.from, string(msg.payload))
		}

	}()

	log.Fatal(server.Start())

	fmt.Println("Hi")
}
