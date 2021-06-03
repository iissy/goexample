package main

import "fmt"

func main() {
	server := new(Server)
	obj := new(InterImpl)
	server.RegisterTransferImpl(obj)
	obj.Msg = "asdf"

	fmt.Printf("%+v\n", obj)
	fmt.Printf("%+v", server.InterIm)
}

type Inter interface {
	Handle()
}

type InterImpl struct {
	Msg string
}

func (i *InterImpl) Handle() {
	fmt.Println("b")
}

type Server struct {
	InterIm Inter
}

func (s *Server) RegisterTransferImpl(inter Inter) {
	s.InterIm = inter
}
