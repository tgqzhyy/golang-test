package main

func main() {
	messages :=make(chan string)
	signals :=make(chan bool)

	select {
	case msg :=<-messages:

	}
}
