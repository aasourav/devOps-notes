package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// variables
var seatingCapacity = 10
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
}

func (shop *BarberShop) addBarber(barber string) {
	shop.NumberOfBarbers++
	go func() {
		isSleeping := false
		color.Yellow("%s goes to the waiting room to check for clients. ", barber)
		for {
			// if there are no clients, the barber goes to sleep
			if len(shop.ClientsChan) == 0 {
				color.Yellow("There is nothing to do, so %s takes a nap.", barber)
				isSleeping = true
			}
			client, shopOpen := <-shop.ClientsChan

			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes %s up.", client, barber)
					isSleeping = false
				}
				// cut hair
				shop.cutHair(barber, client)
			} else {
				// shop is closed, so send the barber home and close the goroutine
				shop.sendBarberHome(barber)
				return
			}
		}
	}()
}

func (shop *BarberShop) cutHair(barber, client string) {
	color.Green("%s is cutting %s's hair.", barber, client)
	time.Sleep(shop.HairCutDuration)
	color.Green("%s is finished cutting %s's hair.", barber, client)
}

func (shop *BarberShop) sendBarberHome(barber string) {
	color.Cyan("%s is going home", barber)
	shop.BarbersDoneChan <- true
}

func (shop *BarberShop) closeSHopForDay() {
	color.Cyan("Closing shop for the day.")
	close(shop.ClientsChan)
	shop.Open = false

	// wait for all the barbers done
	for a := 1; a <= shop.NumberOfBarbers; a++ {
		<-shop.BarbersDoneChan
	}
	close(shop.BarbersDoneChan)
	color.Green("---- The barber shop is now closed for the day ----")
}

func (shop *BarberShop) addClient(client string) {
	//print out message

	color.Green("%s client arrives!", client)
	if shop.Open {
		select {
		case shop.ClientsChan <- client:
			color.Yellow("%s takes a seat in the waiting room.", client)
		default:
			color.Red("The waiting room is full, so %s client leaves.", client)
		}
	} else {
		color.Red("The shop is already closed %s client leaves.", client)
	}
}

func main() {
	// seed our random number generator
	rand.Seed(time.Now().UnixNano())

	// print welcome message
	color.Yellow("The sleeping barber problem")
	color.Yellow("---------------------------")

	// create channels if we need any
	clientChan := make(chan string, seatingCapacity)
	doneChannel := make(chan bool)

	// create the barbershop
	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientsChan:     clientChan,
		BarbersDoneChan: doneChannel,
		Open:            true,
	}

	color.Green("The shop is open for the day.")

	// add barbers
	shop.addBarber("Frank")
	// start the barbershop
	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeSHopForDay()
		closed <- true
	}()

	// add clients
	i := 1
	go func() {
		for {
			// get a random number with avg arrival rate
			randomMilSec := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Microsecond * time.Duration(randomMilSec)):
				shop.addClient(fmt.Sprintf("Cleint #%d", i))
				i++
			}
		}
	}()

	// block until the bargershop is closed
	<-closed
	time.Sleep(time.Second * 5)
}
