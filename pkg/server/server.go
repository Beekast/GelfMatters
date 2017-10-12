/*
Copyright 2017 Beekast.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package server

import (
	"os"
	"fmt"
	"io"
	"bufio"
	"net"

	"github.com/Beekast/GelfMatters/pkg/protocol"
)


type Server struct {
	listener	net.Listener
	sender		chan []string
	proto		protocol.Protocol
}

func New(bind string, port string, p protocol.Protocol) *Server {
	l, err := net.Listen("tcp", bind + ":" + port)
	if err != nil {
		fmt.Println(" [!] Error listening:", err.Error())
		os.Exit(1)
	}
	fmt.Println(" [*] Listening on " + bind + ":" + port)

	return &Server{listener: l, proto: p}
}

func (s *Server) SetSender(sender chan []string) {
	s.sender = sender
}

func (s *Server) AcceptConnections() {
	defer s.listener.Close()

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println(" [!] Error at receiving connection:", err.Error())
			os.Exit(1)
		}

		go s.requestCallback(conn)
	}
}

func (s *Server) requestCallback(conn net.Conn) {
	fmt.Println(" [*] New client connection:")

	buf := bufio.NewReader(conn)

	for  {
		data, err := buf.ReadBytes(0)
		if err != nil {
			if err != io.EOF {
				fmt.Println(" [!] Error at reading data:", err.Error())
			}
			break
		}

		l := len(data)

		// Discard the last byte "\0"
		if err := s.proto.Parse(data[:l-1]); err == nil {
			s.sender <- s.proto.Extract()
		}
	}

	conn.Close()
	fmt.Println(" [*] Connection closed")
}
