package main

type Message struct {
	From    string
	Payload string
}

type Server struct {
	msg  chan Message
	none chan struct{}
}

func (s *Server) start() {
free:
	for {
		select {
		case msg := <-s.msg:
			println(msg.Payload)
		case <-s.none:
			break free
		}
	}
}

func main() {
	println("hello")
}
