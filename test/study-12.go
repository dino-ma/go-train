package main

import "fmt"

type USB interface {
	Name() string
	Connecter
}

type Connecter interface {
	Connect()
}

type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("connect:", pc.name)
}

func main() {
	var usb USB
	usb = PhoneConnecter{name: "name"}
	usb.Connect()
	Disconnect(usb)
}

func Disconnect(usb interface{}) {
	//if pc, ok := usb.(PhoneConnecter); ok {
	//	fmt.Println("disconnected", pc.name)
	//	return
	//}
	//fmt.Println("Unkown decive")
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("disconnect:", v.name)
	default:
		fmt.Println("unkown decive")
	}

}
