package main

import "fmt"

//target
type charger interface {
	chargeAppleMobile()
}

// concrete implementation
type apple struct{}

func (a *apple) chargeAppleMobile() {
	fmt.Println("Charging APPLE mobile")
}

//client
type client struct{}

//Adaptee
type android struct{}

func (a *android) chargeAndroidMobile() {
	fmt.Println("Charging ANDROID mobile")
}

type AndroidAdapter struct {
	android *android
}

func (a *AndroidAdapter) chargeAppleMobile() {
	a.android.chargeAndroidMobile()
}

func (c *client) chargeMobile(ch charger) {
	ch.chargeAppleMobile()
}

func main() {
	apple := &apple{}
	client := &client{}
	android := &android{}
	androidAdapter := &AndroidAdapter{
		android: android,
	}
	client.chargeMobile(apple)
	client.chargeMobile(androidAdapter)
}
