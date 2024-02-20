package user

import (
	"auth-service/app/helpers/log"

	"google.golang.org/grpc"
)

func ConnectUserServiceGrpc() UserClient {
	conn, err := grpc.Dial("localhost:6003", grpc.WithInsecure())
	if err != nil {
		log.GrpcLog(log.Error, "user-service", "can't connect to grpc service")
	}
	// defer conn.Close()

	client := NewUserClient(conn)

	return client

}
