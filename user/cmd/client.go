package cmd

//func main() {
//
//	lis, err := net.Listen("tcp", port)
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//	var opts []grpc.ServerOption
//	s := grpc.NewServer(opts...)
//	pb.RegisterUserServiceServer(s, &UserServiceServer{})
//	log.Printf("server listening at %v", lis.Addr())
//	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
//
//	configMongodb(ctx)
//
//	fmt.Printf("Server succesfully started on port :%v", port)
//	c := make(chan os.Signal)
//
//	signal.Notify(c, os.Interrupt)
//	<-c
//	fmt.Println("\nStopping the server...")
//	s.Stop()
//	lis.Close()
//	fmt.Println("Closing MongoDB connection")
//	client.Disconnect(ctx)
//	fmt.Println("Done.")
//}
