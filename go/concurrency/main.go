package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade , pizzasFailed, total int

type PizzaOrder struct{
	pizzaNumber int
	message string
	success bool
}

type Producer struct{
	// channel is more powerful than mutex and wait groups. bcz they allow one go routine to exchange data with another go routine. they can talk to each other.
	data chan PizzaOrder
	quit chan chan error // bcz `func (p *Producer) Close() error {`
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return  <- ch
}

func makePizza(pizzaNumber int)*PizzaOrder{
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas{
		delay := rand.Intn(5) + 1
		fmt.Println("Received an order number: ",pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		}else{
			pizzasMade++
		}
		total++
		fmt.Printf("Making pizza %d. It will take %d seconds..\n",pizzaNumber,delay)
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2{
			msg = fmt.Sprintf("We run out of ingredients for pizza: %d\n",pizzaNumber)
		}else if rnd <=4 {
			msg = fmt.Sprintf("the cook quite after making pizza: %d\n",pizzaNumber)
		}else{
			msg = fmt.Sprintf("Pizza order %d is ready\n",pizzaNumber)
		}
		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message: msg,
			success: success,
		}
		return &p
	}

	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}


func pizzaria(pizzaMaker *Producer){
	// keep track of which pizz we are making
	i := 0
	// run forever or until we receive a quite notification


	// try to make pizzas
	for{
		currentPizza := makePizza(i)
		// try to make a pizza
		// decision
		
	}
}

func main(){
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// print out a message
	color.Cyan("The Pizzaria is open for business.")
	color.Cyan("------------------------------------")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in the background (own go routine)
	go pizzaria(pizzaJob)
	//once you create a channel when you finish with it the golder rule is
	// you must close it

	// create and run consumer

	// print out the ending message
}