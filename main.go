package main

import (
	dltlsdml_com "dltlsdml.com/grpc/dltlsdml.com"
)

func main() {
	//_chan.ChanTest()
	go dltlsdml_com.TestGrpcServer()

	dltlsdml_com.TestGrpcClient()

}
