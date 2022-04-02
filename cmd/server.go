package main

import (
	"net/http"
	"net"
	"encoding/json"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"server/model"
	"server/grpc/service"
	"server/grpc/pb"
)

var courseList = model.NewCourses()

func main() {
	// Start gRPC server
	go startGrpc()

	// Http server
	http.HandleFunc("/course", CourseListHandler)
	http.ListenAndServe(":8000", nil)
}

func startGrpc() {
	listen, err := net.Listen("tcp", "localhost:50056")
	if err != nil {
		panic(err)
	}

	// gRPC server
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	courseService := service.NewCourseGrpcService()
	courseService.Courses = courseList

	// Register the gRPC server and serve
	pb.RegisterCourseServiceServer(grpcServer, courseService)
	grpcServer.Serve(listen)
}

func CourseListHandler(w http.ResponseWriter, r *http.Request) {
	courseJson, _ := json.Marshal(courseList)
	w.Write([]byte(courseJson))
}
