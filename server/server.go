package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"thriftDemo/gen-go/demo"
)

type CalculatorServer struct {
}

func (e *CalculatorServer) Ping() (err error) {
	fmt.Println("ping.....")
	return nil
}

func (e *CalculatorServer) Add(num1 int32, num2 int32) (r int32, err error) {
	fmt.Println("add...")
	return num1 + num2, nil
}
func (e *CalculatorServer) Calculate(logId int32, work *demo.Work) (r int32, err error) {
	fmt.Println("calculate...")
	return 1, nil
}
func (e *CalculatorServer) Zip() (err error) {
	fmt.Println("zip....")
	return nil
}

func main() {
	transport, err := thrift.NewTServerSocket(":9898")
	if err != nil {
		panic(err)
	}

	handler := &CalculatorServer{}
	processor := demo.NewCalculatorProcessor(handler)

	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTCompactProtocolFactory()
	server := thrift.NewTSimpleServer4(
		processor,
		transport,
		transportFactory,
		protocolFactory,
	)

	if err := server.Serve(); err != nil {
		panic(err)
	}
}
