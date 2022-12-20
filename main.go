package main

import (
	"context"
	"fmt"
	"log"

	m "github.com/shadowshot-x/ActuatorBuf-grpc/protobufs"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9091", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := m.NewActuatorClient(conn)

	fmt.Println(client)

	check, err := client.ContractStateCheck(context.Background(), &m.ContractVariableState{
		Var1: 0,
		Var2: "string123123",
	})

	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	fmt.Println(check)
	fmt.Println("golang")
}
