package main

import (
	"fmt"
)

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

type Huawei struct {

}

func (huawei Huawei)call()  {
	fmt.Println("这是华为")
}
func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

	phone =new(Huawei)
	phone.call()


}
/**
I am Nokia, I can call you!
I am iPhone, I can call you!
这是华为



*/
