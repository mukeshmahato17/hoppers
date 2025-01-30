package main

func main() {
	server := APIServer{listenAddr: ":3000"}
	server.Run()
}
