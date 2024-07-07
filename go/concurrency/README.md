A golder rule for concurrency: If you don't need it, don't use it.

Keep your application's complexity to an absolute minimum; it's easier to write, easier to understand, and eaiser to maintain.

GO's philosophy: Don't communicate by sharing memory, share memory by communicating.

Go Extensions (for vs code):

- Go (official one)

Now For Go environment:
Shift + ctrl + p (linux)
Go: install/Update tool (it will show you all items should install for go lang)

- GO template syntax

### Installing Make (build tool)

Install it.

## Go Routines

Running things in the background, or concurrently
simple to use
It's create problems . Several ways to solve those problems.

```go
package main

import "fmt"

func main(){
	fmt.Println("Hello")
}
```

The main function itself is a go routine.

They run on lightweight threads , not a builtin hardware threads of processor. it handle by the go, they take very little memory , they run very quickly , and they're all managed as group of go routines is called go routines.

They all managed by go schedulers.

```go
package main

import "fmt"

func printSomething(s string){
	fmt.Println(s)
}

func main(){
	// fmt.Println("Hello")
	printSomething("Print this 0")
	printSomething("Print this 1")
	printSomething("Print this 2")


}

```

The code give up. It's run synchronously.

```go
package main

import "fmt"

func printSomething(s string){
	fmt.Println(s)
}

func main(){
	// fmt.Println("Hello")
	go printSomething("Print this 0")
	printSomething("Print this 1")
	printSomething("Print this 2")


}

```

Now if we run this(top). you can see first that I prefixed as `go` will not printed in console.

What happen there. This programme executed so quickly . so that it will not wait for this go routine.

How can I handle this. use sleep (worst solution)

```go
package main

import "fmt"

func printSomething(s string){
	fmt.Println(s)
}

func main(){
	// fmt.Println("Hello")
	go printSomething("Print this 0")
    time.Sleep(1 * time.Second)
	printSomething("Print this 1")
	printSomething("Print this 2")


}

```

So, what's the best solution.
Use wait group.

```go
package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup){
	defer wg.Done() // this will decrement
	fmt.Println(s)
}

func main(){
	// fmt.Println("Hello")
	var wg sync.WaitGroup

	// first we have to add one entry to wait group
	// this is an int. one entry for every thing
	// you have to wait for



	words:= []string{
		"alpha",
		"beta",
		"gamma",
		"peta",
		"hexa",
	}

	wg.Add(5) // why 5 ? bcz words have 5 elements we need to wait for (print)

	for i, x := range words{
		go printSomething(fmt.Sprintf("%d: %s\n",i,x),&wg)
	}

	wg.Wait() // this wait the wait value is set to zero

	wg.Add(1) // if i did not add this the code below will be caused as error
	// the error is : sync: negative WaitGroup counter. so why I write this. wait group is easy to use. be be careful is will not go to the below zero.


	printSomething("this is the second thing to be printed ",&wg)

	// another things is you have no gaurantee which order will print . it's decide by the go.
}

```

Le'ts say you add wait value bigger than You should wait. there will occur deadlock error. so be careful.

### Writing Tests with wait groups

```go
package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSomething(t *testing.T){
	stdOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout  = w
	var wg sync.WaitGroup
	wg.Add(1)

	go printSomething("epsilon", &wg)
	wg.Wait()

	_ = w.Close()

	res, _ := io.ReadAll(r)
	output := string(res)
	os.Stdout = stdOut

	if ! strings.Contains(output,"epsilon"){
		t.Error("expected to find epsilon, bit it's not there")
	}
}

```

## Racing Conditions, Mutexes , and Channels

`sync.Mutex` -> allows us to deal with race conditions.
it's easy to use.
Deal with shared reasources and concurrent/ parallel goroutines.
lock/ unlock
We can test for race conditions when running code, or testing it.

Race conditions occur when multiple GoRoutines try to access the same data.
Can be difficult to spot when reading code.
Go allows us to check for them when running a program , or when testing our code with go test.

### Channels

channels are means of having GoRoutines share data. They can talk to each other . This is go's philosopy of having things share memory be communicating, rather than communicating by sharing memory.

The Producer/Consumer Problem. ( we will solve this using channels)

![alt text](image.png)

Creating Race Condition scenerio:

```go
package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string, wg *sync.WaitGroup){
	defer wg.Done() // this will decrement
	msg = s
}

func main(){
	msg = "hello world"
	wg.Add(2)

	go updateMessage("Hello, uni", &wg)
	go updateMessage("Hello, vercel", &wg)

	wg.Wait()
	// here we can't expect which one will print. line 20 and 21. bcz it's decide
	// by the go.

	// if we run go run -race .
	// it will show an Warning "Data Race". Data race takes place when you have
	// concurrent go routines that access the same piece of data (here msg var).
	// bcz we never sure which one is going to finish first.
	// how do we fix this.
	// we can do that by using Mutex (Mutual  Exclusivity)


	fmt.Println(msg)
}

```

### Now we will fix this:

```go
package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string, mutex *sync.Mutex){
	defer wg.Done()

	mutex.Lock() // now we have exclusive access to this var.
	// nobody else can not change that value. until it's done with it (mutex.Unlock)
	msg = s
	mutex.Unlock()
	// these means I am accessing data safely. this is called thread safe operation
}

func main(){
	msg = "hello world"
	var mutex sync.Mutex

	wg.Add(2)

	go updateMessage("Hello, uni", &mutex)
	go updateMessage("Hello, vercel", &mutex)

	wg.Wait()
	fmt.Println(msg)
}

```

### More complex race conditions.

```go
package main

import (
	"fmt"
	"sync"
)


var wg sync.WaitGroup

type Income struct{
	Source string
	Amount int
}


func main(){
	// var for bank balance
	var bankBalance int
	var balance sync.Mutex

	// print out starting values
	fmt.Println("Initial Balance: ", bankBalance)

	// define weekly revenue
	incomes := []Income{
		{
			Source: "Main Job",
			Amount: 100,
		},
		{
			Source: "Gift",
			Amount: 10,
		},
		{
			Source: "Part time Job",
			Amount: 50,
		},
		{
			Source: "Investments",
			Amount: 70,
		},
	}

	wg.Add(len(incomes))

	// loop through 52 weeks and print out how much is made; keep a running total
	for i , income := range incomes{
		go func(i int, income Income){
			defer wg.Done()
			for week := 1 ; week <= 52 ; week++{
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()
				fmt.Printf("One week %d, you earned %d from %s\n", week,income.Amount, income.Source)
			}
		}(i,income)
	}

	wg.Wait()

	// print out final balance
	fmt.Printf("Final Bank balance %d\n",bankBalance)
}
```

If we run this `go run -race .` then we will not show the warning. we solved the race problem. bcz of lock and unlock.

### Now test this code:

```go
package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T){
   stdOut := os.Stdout
   r, w, _ := os.Pipe()

   os.Stdout = w

   main()

   _ = w.Close()

   res , _ := io.ReadAll(r)
   output := string(res)

   os.Stdout = stdOut

    if !strings.Contains(output, "11960"){
		t.Error("Wrong balance returned")
	}
}
```

### Producer / Consumer problem solve

First read the doc what is it.
