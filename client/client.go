package main

import (
	"context"
	"demo/gen-go/demo"
	"fmt"
	"log"
	"net"
	"os"

	"git.apache.org/thrift.git/lib/go/thrift"
)

func main() {
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTCompactProtocolFactory()

	transport, err := thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "9898"))
	if err != nil {
		_, err = fmt.Fprintln(os.Stderr, "error resolving address:", err)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		os.Exit(1)
	}

	useTransport, _ := transportFactory.GetTransport(transport)
	client := demo.NewCalculatorClient(thrift.NewTStandardClient(protocolFactory.GetProtocol(useTransport), protocolFactory.GetProtocol(transport)))
	if err := transport.Open(); err != nil {
		_, err := fmt.Fprintln(os.Stderr, "Error opening socket to 127.0.0.1:9898", " ", err)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		os.Exit(1)
	}
	defer transport.Close()

	res, err := client.Add(context.Background(), 1, 2)
	if err != nil {
		log.Println("Echo failed:", err)
		return
	}

	log.Println("response:", res)
}
