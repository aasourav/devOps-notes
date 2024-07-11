# Concurrency

Go routine is user-space threads.

concurrency is a composition of independent execution computations, which may or may not run in parallel.
concurrency enables parallelism

![concurrency single thread](./images/image.png)

In this way, multiple processes are sharing CPU

# Parallelism

Parallelism is the ability to execute multiple computations simultaneously.

![parallelism](./images/image1.png)

Concurrency enables parallelism (Need to know details about it.)

### Concurrency in Summary.

Why we need to think about Concurrency?

- In order to run faster, application needs to be divided into multiple independent units and run them in parallel.

## why ther awas a need to build concurrency primitives in Go?

OS => the job of os is to give fair chance for all process access to CPU, memory and other resources.
Process => An instance of a running program is called a process. Process provides environment for program to execute.
When the program executade the os creates a process and allocates memmory in the virtual address space.
the virtual address space will contain Code segments which is compiled machine code .
There is a Data region which contains global variable.  
Heap Segment used for dynamic memory allocation. stack is used for local varibles of function.

![environment for programe to execute](./images/image2.png)

Threads => are smalles unit of execution that CPU accepts. each process has atleast 1 thread. that is main thread. a process can have multiple threads. threads share same address space. each thread has it's own stack. thread can run independent of each other. THe OS scheduler makes scheduling decisions at thread level, not process level. Threads can run concurrently, with each thread taking turn on the individual core, or they can run in parallel with each thread running at the same time on different cores.

![Thread overview](./images/image3.png)

### Thread States:

When the process is created, the main thread is put into the ready queue.

It is in the runnable state.

Once the CPU is available, the thread starts to execute and each thread is given a time slice.

If that time slice expires, then the thread gets preempted and placed back onto the queue.

If the thread gets blocked due to an I/O operation like read/write on disc or network operation

or waiting for the event from other processes, then it is placed in the waiting queue until the I/O

operation is complete.

Once it is complete, it is placed back onto the ready queue.

![Thread states ](./images/image4.png)

### Can we divide our application into Processes and Threads and achive concurrency? => Yes , but there are limitation.

Wy limitations?

- Context switching.

Context switches are considered expensive. CPU has to spend time copying the context of the current executing thread into memory and restoring the context of the next chosen thread And it does take thousands of CPU instructions to do context switching, and it is a wasted time as CPU is not running your application, but doing context switching.

![alt text](./images/image5.png)

In this diagram, you might see the context switching between the threads of the same process is relatively. cheap compared to the context switching between the threads of different processes.

can we scale the number of threads per process?
not much actually. If we scale the number of threads in a process too high, then we hit C10k problem.

what is c10k problem:
the scheduler allocates a time slice for each process to execute on CPU core.

This CPU time slice is divided equally among threds.
![c10k](./images/image6.png)

![alt text](./images/image7.png)

So if we define scheduler period as 10 milliseconds and we have two threads, then each thread is

going to get five milliseconds separately for execution.

If we have five threads, then each thread is going to get two milliseconds to execute.

But what if there are a thousand threads? => 10 microsec
This would be bad, as CPU will be spending more time in context switching than running the application.

So to do any meaningful job, a thread needs at least a minimum of two milliseconds.

![\c10k 1](./images/image8.png)

If a minimum time for the thread is two milliseconds and accordingly we said this scheduler period then

to execute a thousand threads, the scheduler will take two seconds.

If there are 10000 threads, then it will take 20 seconds to complete one cycle of the execution, each

thread will have to wait for 20 seconds for its next execution.

So the application is going to become less responsive.

### So other issue is the stack size, the operating system gives a fixed stack size for each thread,

![stack size issue](image9.png)

the actual size depends on the hardware.

On my machine, it is 8MB.

So if I have a 8GB of memory, then in theory I can only create 1000 threads.

So the fixed stack size limits the number of threads that we can create to the amount of memory we have.

#### let us summarize.

We saw what is a process, a process is an instance of a running program, and it provides an environment

for the program to execute.

We saw what is a thread, a thread is the smallest unit of execution, and every process has atleast

one thread and process can have multiple threads and all threads share the same address space.

And we saw what are the limitations with threads?

Fixed stack size, fixed stack size limits, the number of threads that we can create to the amount of memory

we have.

and C10k problem, as we scale the number of threads the scheduler cycle is going to increase and the application

### WHy Concurrency is hard?

In this module, we will see why concurrency is hard and how sharing of memory between the threads can

create a lot of complexity.

If you remember from the previous module, you know that all the threads share the same address space.

They share the heap and the data region of the process.

and threads communicate between each other by sharing memory.

But this sharing of memory creates a lot of complexity with concurrently executing threads.

So if two threads are running concurrently and they try to access the same area of memory with one thread

trying to write to the memory, then there will be a data race and the outcome of the program will be

un-deterministic.

Let us consider an example with thread 1 and thread 2 are running concurrently.

and here they are trying to increment the value of a global variable i.

![alt text](image.png)

The increment operation is not atomic, at the code level

it looks as a one statement, but in the context of the machine instructions, it involves retrieving

the value of i from memory, incrementing the value of i and storing the value of i to the memory.

So what happens if the thread gets preempted between these operations?

So let us see some scenarios.

Now, let us take two sequence of execution, in the first sequence, thread 1 and thread 2 are executing

![alt text](image-1.png)

sequentially, one after another.

In second sequence, thread 1 and thread 2 are executed in interleaved fashion, where execution of one

thread is preempted by the other thread.

Let us consider the first scenario, thread 1 starts the execution, it retrieves the value of i, which

will be zero, it increments it and then stores the value of i.

![alt text](image-2.png)

then thread 2 to come along, it retrieves the value of i, which will be one, and it increments by one and

it stores the value of i, which will be 2.

![alt text](image-3.png)

This is fine, now let us consider the second scenario.

Thread 1 starts the execution, it retrieves the value of i, which will be zero, then it increments

the value of i to one.

But before thread 1 can write the value of i to memory,

it gets preempted and thread 2 starts the execution.

it retrieves the value of i, which will be zero,

then it increments it

and it stores the value of i as 1.

![alt text](image-4.png)

Now, thread 1 comes along and it will store the value of i as 1.

![alt text](image-5.png)

As you see, the value of i, can be 2 or it can be 1, depending on how the threads are executing.

![alt text](image-6.png)

So concurrent access to memory, leads to un-deterministic outcomes.

So one way to handle this will be to use memory access synchronization tools.

We need to guard the access to the shared memory so that a thread has exclusive access at a time.

and we need to force thread 1 and thread 2 to run sequentially, to increment the value of i.

We can do this by putting a lock on the increment operation.

![alt text](image-7.png)

Putting a lock around a shared memory is a developer's convention, so any time a developer wants to

access the shared memory, they need to acquire the lock and when they are done, they need to release

the lock.

If the developer does not follow this convention, then we have no guarantee of exclusive access and

it can always happen that some code can always sneak in, which does not follow this convention.

and we hitting the race condition at some time.

There are other problems like locking actually reduces the parallelism.

as locks force the threads to be executed sequentially.

So the critical section where we access the shared memory becomes a bottleneck between the threads.

The other problem is, coding mistakes, in-appropriate use of locks can lead to deadlocks.

Let us see an example.

![alt text](image-8.png)

Here there are two threads executing concurrently.

Thread1 starts the execution, it takes a lock on resource v1, it gets preempted by thread2

it takes a lock resource v2

Now, thread1 comes along, it wants a lock on resource v2, but it is not available,

so it goes into the waiting state.

Now, thread2 comes along and it wants a lock on resource v1, but it is not available, so it also goes into

waiting state.

![alt text](image-9.png)

So as you see, this is a circular wait, which leads to deadlock and the application will just hang.

So we have come to the end of this module, let us summarize, so we saw why concurrency is hard.

Sharing of memory between the threads creates complexity.

and concurrent access to the shared memory can lead to race conditions and outcome can be un-deterministic.

Memory access synchronization tools actually reduces the parallelism and comes with its own limitations.

### GoRoutines

In the previous module, we had seen that, there are limitations with threads, the actual number of threads

that we can create is limited, and sharing of memory leads to a lot of complexity with concurrently

executing threads.

In this module, we will see how Go implements concurrency and how Go overcomes some of the

limitations with threads.

Concurrency in Go is based on the paper written by Tony Hoare, communicating sequential processes

or CSP.

The beauty of CSP is that it is very simple, it is based on three core ideas.

Each process is built for sequential execution.

Every process has a local state and the process operates on that local state.

If we have to transfer data from one process to another process, we do not share memory, but we communicate

the data, we send a copy of the data over to other process.

Since there is no sharing of memory, there would be no race, condition or deadlocks,

and we can scale easily, as each process can run independently.

If the computation is taking more time, we can add more processes of the same type and run the computation

faster.

So what tools Go provides for concurrency?

- Goroutines - goroutines are concurrently executing functions,
- Channels - channels are used to communicate data between the goroutines.
- Select - Select is used to multiplex the channels,
- Sync package - Sync package provides classical synchronization tools like the mutex, conditional variables and others.

Goroutines are user space threads, managed by Go runtime, Go runtime is part of the executable, it is built into the executable of the application.Goroutines are extremely lightweight, goroutines starts with 2KB of stack, which can grow and shrink as required.

It has a very low CPU overhead, the amount of CPU instructions required to create a goroutine is very less. This enables us to create hundreds of thousands of goroutines in the same address space. The data is communicated between the goroutines using channels, so sharing of memory can be avoided.

The context switching is much cheaper than the thread context switching as goroutines how less state to store.
Go runtime can be more selective in what data is persisted, how it is persisted and when persisting needs to occur.

Go runtime creates OS threads, goroutines runs in the context of the OS thread.This is important.

![alt text](image-10.png)

Goroutines are running in the context of the OS threads.
Many goroutines can execute in the context of the single OS thread. The operating system schedules the OS threads and the Go runtime schedules, multiple goroutines on the OS thread.

![alt text](image-11.png)

For the operating system, nothing has changed, it is still scheduling the threads, as it was. Go runtime manages the scheduling of the goroutines on the OS threads.

### Summarize GoRoutines

What are Goroutines?
Goroutines are userspace threads managed by go runtime.

We saw what are the advantages of Goroutines over OS threads.

Goroutines are extremely lightweight as compared to OS threads, they start with a very small stack size of 2KB as opposed to 8MB of stack size for the OS threads.

The context switching is very cheap as it happens in the userspace.Goroutines have a very less state to be stored.

This enables us to create hundreds of thousands of goroutines in the same address space.

## Coding exercise link:

git clone https://github.com/andcloudio/go-concurrency-exercises.git

Now see the exercise from the git:

### Solution can be found in : go-concurrency-exercise/01-exercise-solution

### go-concurrency-exercise/01-exercise/01-goroutines/01-hello

### go-concurrency-exercise/01-exercise/01-goroutines/02-client-server/

go routines in the server to handle multiple conrcurrent client connections

## WaitGroups (sunc.WaitGroup)

One of the complexity to manage in concurrency is race condition, the race condition occurs when two or more operations must be executed in the correct order to produce the desired result.

But the program has not been written so that that order is guaranteed to be maintained.

Most of the time, the race condition is introduced due to developers thinking that the program executes in the order they are coded.

In concurrent programs, that is not the case.

In this code snippet, there is no guarantee that the goroutine will be created and executed before main routine checks the value of the data.

#### Goroutines are executed asynchronously from the main routine.

The order in which the main routine and the goroutine will execute is undeterministic.

In this example, there are three possible outcomes.

![alt text](image-12.png)

The first outcome can be nothing is printed, if the sequence of execution is, goroutine gets scheduled before main routine checks the value of the data.

So in this case, the value of the data will be 1, so nothing will be printed.

The second possible outcome could be it will present value is zero.

If goroutine does not get executed before main routine, checks the value of the data.

main routine, checks the value of the data, which will still be zero, so it prints the output as value is zero.

The third possible outcome can be, it print the value is one.

If goroutine gets scheduled between main routine, checking the value of data and printing the value of data, main routine checks the value of data, which will be zero.

So it drops into the if block.

But before executing the print statement, goroutine gets scheduled and increment the value of the data.

And then main routine, execute the print statement. Then the output will be value is one.

Can we bring some determinism into our program?

Can we make main routine, wait for the goroutine to execute before checking the value of data?

This is where sync.waitGroup comes into play.

![alt text](image-13.png)

Go follows logical concurrency model called fork and join.

Go statement forks a goroutine, when a goroutine is done with its job, it joins back to the main routine.

If main does not wait for the goroutine, then it is very much possible that the program will finish before goroutine gets a chance to run.

In order to create a join point, we use sync.WaitGroup, to deterministically block the main routine.

We create a variable of type sync.WaitGroup we call the add method to indicate the number of goroutines that we are creating.

we call done method inside the goroutine closure (defer wg.Done()(in image give top)) to indicate that the goroutine is exiting.

Usually it is called with the `defer`, to make sure that it gets called on all the exit points of the function.

Wait method will block the main routine until all the goroutines have exited.

wait group is like a concurrent counter

call to add increases the counter by the integer that is passed in, call to done will decrement the counter by one, call to wait, will block until the counter becomes zero.

Here you notice (image give top) that add method is placed outside the goroutine enclosure had to call to add was placed inside the goroutine closure, then it is very much possible that the call to wait method could execute before the add method.

In that case, the main routine could return without blocking.

Therefore, we call the add method outside the goroutine enclosure to make sure that the add method gets called before wait.

#### Exercise WaitGroup.

### go-concurrency-exercise/01-exercise/01-goroutines/03-join

## GoRoutines and Clousers

Goroutines executes within the same address space they created in.

They can directly modify the variables in the enclosing lexical block.

This enables developers to write goroutines within the <b>lexical block</b> of the enclosing function as a sequential code. Go compiler and the runtime takes care of pinning the variable, moving the variable from stack to heap, to facilitate goroutines, to have access to the variable even after the enclosing function has returned.

![alt text](image-14.png)

### go-concurrency-exercise -> 01-exercise -> 01-goroutines -> 05-closure

Explain the solution code:

We have an exercise here, we need to run the program and check that variable i was pinned for access from goroutine even after the closing function returns.

So let us see what we have here. We have a function, inside function we have a local variable(`var i int`) . We are spinning a goroutine(`go func()`) and we are returning from the function.

Inside the goroutine, we are accessing the local variable of the function and we are incrementing its value and we are printing the value.

In the main routine we are calling a function and we are waiting for the goroutines to execute.

let me run this program.

So what's happening here is the function has returned, but goroutine still has the access to the local variable of the function.

So usually when the function returns, the local variables goes out of scope.

But here, the runtime is clever enough to see that the reference to a local variable i is still being held by the goroutine,so it pins it, it moves it from the stack to heap, so that goroutine still has the access to the variable even after the enclosing function returns

### go-concurrency-exercise -> 01-exercise -> 01-goroutines -> 06-closure

We have been asked to predict the output of this program and we need to find the issue and fix the issue.

We have a loop, and in each iteration of the loop, we are spinning a goroutine and inside the goroutine, we are printing the value of the interator i and in the main routine, we are waiting for the goroutines to execute.

So what do you think would each goroutine will print the value of i as? do you think they are going to print it as one, two and three respectively or something else?

Let's execute this.

They are all printing the value of i as four, this is because by the time goroutine got the chance to run the value of i had already been incremented to value four.

Now, how can we fix that?

We need to pass the value of i as a parameter to the goroutine function so that goroutine operates on the input that has been passed to it.

So we need to parse the value of i as a parameter here, now let us try to execute.

Now it's going to print one, two, three, here the order depends on the order of the execution of the goroutines.

So in this exercise, what we learnt was goroutines operate on the current value of the variable at

the time of their execution.

If we want the goroutines to operate on a specific value, then we need to pass that as an input to the goroutine.

## Deep Dive Go Scheduler.

In this module and in the next couple of modules, we will have a deep dive into Go scheduler and see how Go scheduler works underneath.

Go scheduler is part of the Go runtime, Go runtime is part of the executable, it is built into the executable of the application.

Go scheduler is also known as M:N scheduler.It runs in the user space.Goroutines are scheduled on the OS threads by the Go scheduler.

So a goroutine runs in the context of the OS thread.

Go runtime creates a number of worker OS threads, equal to GOMAXPROCS environment variable value, the default value is the number of processors on the machine. So if we have 4 cores, then 4 OS threads will be created.

If you have 8 cores, then 8 OS threads will be created.

It is the responsibility of the Go scheduler to distribute runnable goroutines over multiple threads that are created.

At any time N goroutines could be scheduled on M OS threads that runs on at most GOMAXPROCS number of processors.

As of Go 1.14, the Go scheduler implements <b>Asynchronous Preemption.</b>

It used to be co-operative scheduler, but then the question came, what happens when a long running routine just hogs onto the CPU?

other goroutine would just get blocked.

So that's the reason why asynchronous pre-emption was implemented.

So in asynchronous preemption, what happens is, a goroutine is given a time slice of ten milliseconds for execution.

When that time slice is over, Go scheduler will try to preempt it, this provides other goroutines the opportunity to run even when there are long running CPU bound goroutines scheduled.

Similar to threads, goroutines also have states.

![alt text](image-15.png)

When it is created, it will be in runnable state, waiting in the run queue.

It moves to the executing state once the goroutine is scheduled on the OS thread.If the goroutine runs through its time twice, then it is preempted and placed back into the run queue.

If the goroutine gets blocked on any condition, like blocked on channel, blocked on a syscall or waiting for the mutex lock, then they are moved to waiting state.

Once the I/O operation is complete, they are moved back to the runnable state.

Now we will look into different elements involved in Go scheduling.

![alt text](image-16.png)

For a CPU core, Go runtime creates a OS thread, which is represented by the letter M. OS thread works pretty much like POSIX thread. Go runtime also creates a logical processor P, and associate that with the OS thread M.

The logical processor holds the context for scheduling, which can be seen as a local scheduler running on a thread.

G represents a goroutine running on the OS thread.

Each logical processor P has a local run queue where runnable goroutines are queued. Here it is depicted by the colored circles.

There is a global run queue, once the local queue is exhausted, the logical processor will pull goroutines from global run queue.

When new goroutines are created, they're added to the end of the global run queue.

Let's see a context switch.

![alt text](image-17.png)

Goroutine G1 has reached a scheduling point,

then the logical processor will pop a goroutine from its local run queue in this case G2 and sets the stack and the instruction pointer for the goroutine G2 and begins running that goroutine, the previously running goroutine G1, is placed back into the local run queue.

![alt text](image-18.png)

As you see, there is no change as far as the OS is concerned.

It is still scheduling the same OS thread.

The context switching between the goroutines is managed by the logical processor.

There is a one to one mapping between OS thread and the logical processor, if there are two cores and we have set GOMAXPROC environment variable to 2, then go runtime, creates another OS thread and logical processor, and associates the OS thread with the logical processor, and goroutines can be scheduled on the second OS thread.

![alt text](image-19.png)

Let us summarize.

We saw how Go scheduler works.

Go runtime has a mechanism known as M:N scheduler, where N goroutines could be scheduled on M OS threads that run on at most GOMAXPROC number of processors.

As of Go 1.14 Go scheduler implement asynchronous preemption where each goroutine is given a time slice of ten milliseconds for execution.

We saw, what are the components of Go scheduler. OS thread is represented by the letter M.

P is the logical processor which manages scheduling of goroutines. G is the goroutine, which includes the scheduling information like stack and instructions pointer. Local run queue is where runnable. goroutines are queued.

When a goroutine is created, they are placed into the global run queue.

## Deep Dive Context switching.

In this module we will see context switching caused due to synchronous system call.

What happens in general when a goroutine makes a synchronous system call, like reading or writing to a file with sync flag set.

There will be a disc I/O to be performed, so synchronous system call will block for I/O operation to complete.

Due to which the OS thread can be moved out of the CPU and placed in the waiting queue for the disc I/O to complete.

So we will not be able to schedule any other goroutine on that thread.

The implication is that synchronous system call can reduce parallelism.

So how does Go scheduler handle this scenario?

Let us see.

![alt text](image-20.png)

Here goroutine G1 is running on OS thread M1.

G1 is going to make synchronous system call, like reading on a file, that will make the OS thread M1 to block.

Go scheduler identifies that G1 has caused OS thread M1 to block, so it brings in a new OS thread, either from the thread pool cache or it creates a new OS thread if a thread is not available in the thread pool cache.

![alt text](image-21.png)

Then Go scheduler will detach the logical processor P from the OS thread M1, and moves it to the new OS thread M2.

![alt text](image-22.png)

G1 is still attached to the old OS thread M1.

The logical processor P can now schedule other goroutines in its local run queue for execution on the OS thread M2.

Once the synchronous system call that was made by G1 is complete, then it is moved back to the end of the local run queue on the logical processor P.

![alt text](image-23.png)

And M1 is put to sleep and placed in the thread pool cache.

So that it can be utilized in the future when the same scenario needs to happen again.

So let us summarize, we saw, how context switching works when a goroutine calls synchronous system call.

When a goroutine makes a synchronous system call, Go scheduler brings new OS thread from thread pool cache. And it moves the logical processor to the new thread. Goroutine that made the system call, will still be attached to the old thread.other goroutines in the local run queue are scheduled for execution on the new thread.

Once the system call returns, the goroutine which made the system call, is moved back to the local run queue of the logical processor and old thread is put to sleep.

## Context switching due to Asynchronous calls

In this module, we will look into context switching due to a asynchronous system calls, like the network system call or http api call.

What happens when a asynchronized system call is made?

Asynchronous system call happens when the file descriptor that is used for doing network I/O operation is set to non-blocking mode.

If the file descriptor is not ready, for example, if the socket buffer is empty and we are trying to read from it, or if the socket buffer is full and we are trying to write to it, then the read or the write operation does not block, but returns an error.

And the application will have to retry the operation again at a later point in time.

So this is good, but it does increases the application complexity.

The application will have to create any event loop and set up callbacks, or it has to maintain a table

mapping the file descriptor and the function pointer, and it has to maintain a state to keep track of how much data was read last time or how much data was written last time. And all these things, does add up to the complexity of the application.And if it is not implemented properly, then it does make the application a bit inefficient.

So how does Go handle this scenario?

Go uses <b>netpoller.</b>

There is an abstraction built in syscall package.

syscall package uses netpoller to convert asynchronous system call to blocking system call. when a goroutine makes an asynchronized system call, and file descriptor is not ready, then the Go scheduler uses netpoller OS thread to park that goroutine.

The netpoller uses the interface provided by the operating system, like epoll on Linux, kqueue on MacOS, iocp on Windows, to poll on the file descriptor.

Once the netpoller gets a notification from the operating system, it in-turn notifies the goroutine to retry the I/O operation.

In this way, the complexity of managing asynchronous system call is moved from the application to go runtime.

So the application may not have to make a call to select or poll and wait for the file descriptor to be ready, but instead it will be done by the netpoller in an efficient manner.

Let us look into an example.

Here G1 is executing on the OS thread M1.

![alt text](image-24.png)

G1 opens an network connection with net.Dial

The file descriptor used for the connection is set to non-blocking mode.

When the goroutine tries to read or write to the connection, the networking code will do the operation until it receives an error.

EAGAIN

Then it calls into the netpoller, then the scheduler will move the goroutine G1 out of the OS thread M1 to the netpoller thread.

And another goroutine in the local run queue, in this case G2 gets scheduled to run on the OS thread M1.

![alt text](image-25.png)

The netpoller uses the interface provided by the operating system to poll on the file descriptor.

When the netpoller receives the notification from the operating system that it can perform an I/O operation on the file descriptor, then it will look through its internal data structure.

To see if there are any goroutines that are blocked on that file descriptor.

Then it notifies that goroutine, then that goroutine can retry the I/O operation. Once the I/O operation is complete, the goroutine is moved back to the local run queue and it will be processed, by the OS thread M1 when it gets a chance to run.

![alt text](image-26.png)

In this way to process an asynchronous system call, no extra OS thread is used, instead the netpoller OS thread is used to process the Go routines.

So let us summarize.

So in this module, we saw what happens when a goroutine makes a asynchronous system call.

Go uses netpoller to handle asynchronous system call. netpoller uses the interface provided by the operating system to poll on the file descriptor.

And it notifies the Goroutine to try the I/O operation when it is ready.

In this way, the application complexity of managing an asynchrous system call is moved to the

Go runtime, which manages it in an efficient manner.

## Work Stealing.

In this module, we will look into work stealing concept in Go scheduler.Work stealing helps to balance the goroutines across the logical processors.So that work gets better distributed and gets done more efficiently.

Let us look into an example, here we have a multithreaded go program, we have 2 OS threads and 2 logical processors, the goroutines are distributed among the logical processors.

![alt text](image-27.png)

Now, what happens if one of the logical processor services all its goroutines quickly? the global run queue and local run queue of P2.

We see that P1 has no more goroutines to execute, but there are goroutines in runnable state in

![alt text](image-28.png)

The work stealing rule says that, if there are no goroutines in the local run queue, then try to steal from other logical processors.

If not found, check the global run queue for the goroutines.

If not found, check the netpoller.

![alt text](image-29.png)

In this case, P1 does not have any runnable goroutine in its local run queue, so it randomly picks another logical processor, P2 in this case and steals half of its goroutines from its local run queue.

![alt text](image-30.png)

We see P1 has picked up goroutines, G7 and G8 to its own local run queue.And P1 will be able to execute those goroutines.Now we are able to better utilize the CPU cores and the work is fairly distributed between multiple logical processors.

![alt text](image-31.png)

What happens when P2 finishes executing all its goroutines?

And P1 one does not have any goroutine in its local run queue.

Then, according to work stealing rule, P2 will look into the global run queue and finds goroutine G9.

![alt text](image-32.png)

![alt text](image-33.png)

G9 get scheduled on OS thread M2.

Let us summarize

In this module, we saw how work stealing scheduler works.

If the logical processor runs out of goroutines in its local run queue, then it will steal goroutines from other logical processors or global run queue.

So, work stealing helps to balance goroutines across the logical processor and work gets better distributed and gets done more efficiently.

## Channels

![alt text](image-63.png)

Here we have a code snippet, where goroutine is making a computation, and we want to get the result of that computation in our main routine without having to share the memory.

So how can we do that?

This is where channels comes into picture. channels are used to communicate data between the goroutines. channels can also help in synchronizing the execution of the goroutines, one goroutine can let know another goroutine, in what stage of the computation they are in and synchronize their execution.

Channels are typed, they are used to send and receive values of a particular type.
They are thread safe, so the channel variables can be used to send and receive values concurrently by multiple goroutines.

It is very easy to create channels and we declare a variable with chan keyword, followed by the type,

```go
var ch chan T
```

the default value of the channel is nil.
So we need to use built-in function make, to allocate memory for the channel.

```go
var ch chan T
ch = make(chan T)

//or
ch := make(chan T)
```

And the make function returns a reference for the allocated memory.

Or we can use a short variable declaration with make built-in function, which declares and allocates memory for the channel in one statement.

Pointer operators can be used to send and receive values from the channel, and the arrow direction indicates

```go
  // <-
  // send
  ch <-v
  // receive
  v = <-ch

```

the direction of the data flow.

For send, the arrow direction indicates that the value is being written to the channel.

And for receive, the arrow direction indicates that the value is being received from the channel and copied to the variable.

channels are blocking, the sending goroutine is going to block until there is a corresponding receiver goroutine ready to receive the value.

![alt text](image-34.png)

Similarly, the receiver goroutine is going to block until there is a corresponding sender goroutine, sending the value.

And it is the responsibility of the channel to make the goroutine, runnable again once it is ready to receive or send value.

Closing of the channel is very useful for the sender goroutine to indicate to the receiver goroutine, that the sender has no more values to send on the channel and the receiver can unblock and proceed with its other computation.

```go
 close(c)
```

Receive returns two values, the first one is a received value from the channel.

```go
 //receive returns two values
 value, ok = <-chan

```

The second is a boolean value, which indicates whether the value that is being read from the channel is a value that is generated by a write or a default value that is being generated by a close of the channel.

So the second return value will be true if the value is generated by write or it's going to be false, If it is generated by close, and this is very useful to determine whether the value is from write or whether the values from close.

## Exercise Channel -> 01-exercise/02-channel/01-channel

## Range Over the channel

![alt text](image-35.png)

Range over the channel, the receiver goroutine can use range to receive a sequence of values from the channel. range over the channel will iterate over the values received from a channel.

The loop automatically breaks when the channel is closed. So once the sender goroutine has sent all of its values, it will close the channel and the receiver goroutine will break out of the range loop. The range does not return the second boolean value.

Normally the receive returns the second boolean value, but range just returns value, as on close, the range will automatically break out of the loop.

Unbuffered channels:
the channels that we have been creating till now are unbuffered channels.

![alt text](image-36.png)

There is no buffer between the sender goroutine and the receiver goroutine.

Since there is no buffer, the sender goroutine will block until there is a receiver, to receive the value, and the receiver goroutine will block until there is a sender, sending the value.

In buffered channels, there is a buffer between the sender and the receiver goroutine, and we can specify the capacity, that is the buffer size, which indicates the number of elements that can be sent without the receiver being ready to receive the values.

![alt text](image-37.png)

The sender can keep sending the values without blocking, till the buffer gets full, when the buffer gets full, the sender will block.

The receiver can keep receiving the values without blocking till the buffer gets empty, when the buffer gets empty, the receiver will block.

The buffered channels are in-memory FIFO queues, so the element that is sent first, will be the element that will be read first.

### Exercise 01-exercise/02-channel/02-channel

### Exercise 01-exercise/02-channel/03-channel

### Channel Direction

When using channels as functional parameters, you can specify if the channel is meant only to send or only to receive values.

And this specificity will help us to increase the type safety of the programs, in the below example, in is a receive only channel, note the syntax, it's a pointer operator followed by the chan keyword, and out is a send only channel, and the syntax is, chan keyword followed by the pointer operator.

![alt text](image-38.png)

In this example, the pong function can use in, only to receive values. It cannot use this channel to send values. If it tries to send values on this channel, the compiler is going to report an error.so in this way, we can control what operations that function can do with the channels that are passed as parameters.

### Channel direction exercise. 01-exercise/02-channel/04-channel

## Channel ownership

![alt text](image-39.png)

Now we will look into the things that we should be aware when working with channels, and this will help us in troubleshooting.

Default values -
when a channel is declared, its default value is nil.

So we should allocate memory by using the built-in function make.

If that does not happen and we try to send or receive on that channel, then it's going to block forever.

![alt text](image-40.png)

Similarly, closing on the new channel will panic, so we should always make sure that the channels are initialized with the built-in function make.

![alt text](image-41.png)

How we use the channels is important to avoid deadlocks and panics.

We can follow some of the Go idioms.

The best practice is that the goroutine that creates the channel will be the goroutine that will write to the channel and is also responsible for closing the channel.

The goroutine that creates writes and closes the channel is the owner of the channel and the goroutine that utilizes the channel will only read from the channel.

So establishing the ownership of the channel will help us to avoid deadlocks and panics, and it will help in avoiding scenarios like deadlocking by writing to nil channel, closing a nil channel, writing to a closed and closing channel more than once, which can all lead to panic.

### Exercise Channel ownership 01-exercise/02-channel/04-channel

## Deep dive channel

In this module and in the next couple of modules, we will try to understand the mechanics behind channels, how channels work and how to send and receive works underneath.

We use built-in function make to create channels. Here we are creating a buffered channel with three elements.

![alt text](image-42.png)

Internally, the channels are represented by the `hchan` structure. Now let us look into different fields in the hchan struct.

![alt text](image-43.png)

It has a mutex lock field, any goroutine doing any channel operation must first acquire the lock on the channel.buf is a circular ring buffer where the actual data is stored.

And this is used only for the buffered channels, data queue size, is the size of the buffer. qcount indicates a total data elements in the queue. sendx and recvx indicates the current index of the buffer from where it can send data, or receive data.

recvq and sendq are the waiting queues which are used to store blocked goroutines, the goroutines that were blocked while they were trying to send data, or while they were trying to receive data from the channel. waitq, is the linked list of goroutines, the elements in the linked list is represented by the sudog struct.

![alt text](image-44.png)

In the sudog struct, we have the field g, which is a reference to the goroutine, and elem field is pointer to memory, which contains the value to be sent, or to which the received value will be written to.

![alt text](image-45.png)

When we create a channel with built-in function make, hchan struct is allocated in the heap, and make returns a reference to the allocated memory.

And since ch is a pointer, it can be sent between the functions which can perform, send or receive operation on the channel.

![alt text](image-46.png)

This is a runtime values of hchan struct. buf is been allocated a ring buffer and dataq size is set to 3, this value comes from the parameter that has been passed to the make function.

And current qcount is zero, as no data has been enqueued yet.

So in this module, we looked into how channels are represented.

## How send and Receive buffered channel

Let us now look into what happens when we do send or receive on a buffered channel.

![alt text](image-47.png)

In this code snippet, we have 2 goroutines goroutine G1 is sending a sequence of values into the channel, and goroutine G2 is receiving the sequence of values by ranging over the channel. Now, when we create a channel, this will be the representation.

![alt text](image-48.png)

There is a circular queue with size three, which is currently empty.

![alt text](image-49.png)

Let us now consider the scenario when the G1 executes first, G1 is trying to send a value on the channel, which has empty buffer.

First, the goroutine has to acquire the lock on the hchan struct.
![alt text](image-50.png)

Then it enqueues the element into the circular ring buffer.

![alt text](image-51.png)

Note that this is a memory copy. The element is copied into the buffer. Then it increments the value of the sendx to 1. Then it releases the lock on the channel and proceed with its other computation.

![alt text](image-52.png)

Now G2 comes along and tries to receive the value from the channel.

![alt text](image-53.png)

First, it has to acquire the lock on the hchan struct, then it dequeues the element from the buffer queue and copies the value to its variable, v.

![alt text](image-54.png)

And it increments the receive index by 1 and releases the lock on the channel struct and proceeds with its other computation.

![alt text](image-55.png)

![alt text](image-56.png)

This is a simple send and receive an a buffered channel. The points to note are, there is no memory sharing between the goroutines. The goroutines copy elements to and from hchan struct and hchan struct is protected by the mutex lock.

So this is where the Go's tag line comes from. Do not communicate by sharing memory, but instead share memory by communicating.

![alt text](image-57.png)

So in this module we saw a simple send and receive on a buffered channel and the next module will look into what happens when the buffer is full.

## Buffer full Scenerio

Now, let us consider the Buffalo scenario.

![alt text](image-58.png)

G1 enquees the values 1 2 3. buffer gets full and G1 wants to send value 4?

Now, since the buffer is full, what will happen, it will get blocked and it needs to wait for the receiver, right?

Now, how does that happen?

G1 creates `sudog` `G` and G element will hold the reference to the goroutine G1
And the value to be sent will be saved in the elem field.

![alt text](image-59.png)

This structure is enqueed into the `sendq` list.

![alt text](image-60.png)

Then G-1 calls on to the scehduler with call to `gopark()`.

The scheduler will move G1 out of the execution on the OS thread and other goroutine in the local run queue gets scheduled to run on the OS thread.

Now G2 comes along and it tries to receive the value from the channel.

![alt text](image-61.png)

It first, select `look`. deques the element from the Queue. And copies the value into it's variable.

![alt text](image-62.png)

And pops the waiting G1 on the same queue and includes the value saved in the elem field?

That is a value 4 into the buffer. This is important.

It is G2, which will enqueue the value to the buffer on which G1 was blocked.

And this is done for optimization, as G1 late in an have to do any channel operation

again.

Once enqueue is done G2 sets the state of goroutine G1 to runnable.

And this is done by G2 calling `goready(G1)`.
![alt text](image-64.png)

Then G1 is moved to the runnable state and gets added to the local run queue.

![alt text](image-65.png)

And G1 will be scheduled to run on os thread when it gets it chance.

To summarize,
we saw what happens in the case when the giant buffer is full and goroutine tries to send value.

This in the goroutine gets blocked, thisparked on sendq, the data is saved in the elem field of th sudog structure

when Receiver comes along, it dequeues the value from buffer

enqueues the data from elem field to the buffer

and Pops the goroutine in sendq, and puts it into runnable state.

## Buffer empty scenerio

What happens when a goroutine G2 executes first and tries to receive on an empty channel?

![alt text](image-66.png)

The buffer is empty, and G2 has called a receive on an empty channel.
![alt text](image-67.png)

So G2 creates a sudog struct for itself and enqueues it into the receive queue of the channel and the elem field is going to hold the reference to a stack variable v,

![alt text](image-68.png)

And G2 calls upon the scheduler with the call to gopark function, the scheduler will move G2 out of the OS thread and does a context switching to the next goroutine in the local run queue.

![alt text](image-69.png)

Now G1 comes along and tries to send the value on the channel. First, it checks if there are any goroutines waiting in the receive queue of the channel and it finds G2.

![alt text](image-70.png)

Now, G1 copies of the value directly into the variable of the G2 stack and this is important. G1 is directly accessing the stack of G2 and writing to the variable in the G2 stack.

![alt text](image-71.png)

This is the only scenario where one goroutine accesses the stack of another goroutine, and this is done for the performance reasons so that later G2 need not have to come and do one more channel operation and there is one fewer memory copy.

Then G1 pops G2 from the receive queue and puts it into the runnable state,

![alt text](image-72.png)

by calling the go ready functin G2

![alt text](image-73.png)

Now G2 moves back to the local run queue and it will get scheduled on the OS thread M1 when it gets a chance to run.

![alt text](image-74.png)

Now, what we saw here, we saw the buffer empty scenario.

So when goroutine calls receive on an empty buffer, the goroutine is blocked and parked to the receive queue.

The elem field in the sudog struct holds the reference to the stack variable of the receiver goroutine.

The sender goroutine comes along and sender finds the goroutine in the receive queue.

And the sender goroutine copies the data directly into the stack variable of the receiver goroutine.

And pops the receiver goroutine in the receive queue and puts it into the runnable state.

This was about what happens when the receive is called on the empty buffer.

## Send and Receive Unbuffered channels

In this module, we will look into send and receive on an unbuffered channels.

let us see send on unbuffered channel

When the sender goroutine wants to send values on the channel, if there is a corresponding receiver goroutine

waiting in the receive queue, then the sender will write the value directly into the receiver goroutine's stack variable.

The sender routine will then put the receiver goroutine back to the runnable state.

If there is no receiver goroutine in the receive queue, then the sender gets parked into the send queue.

And the data is saved in the elem field in the sudog struct, when the receiver comes along, it copies the data and puts the sender back to the runnable state.

This was about what happens when we do send on an unbuffered channel.

Now, let us see what happens when we do receive on an unbuffered channel.

the receiver goroutine wants to receive value on the channel.

If it finds a sender goroutine in the send queue, then the receiver copies the value in the elem field of the sudog struct to its variable.

Then puts the sender goroutine back to the runnable state. If there was no sender goroutine in the send queue. Then the receiver gets parked into the receive queue.

And a reference to the variable is saved in the elem field in this sudog struct, when the sender comes along,

it copies the data directly into the receiver stack variable.

And puts the receiver back to the runnable state.

So this is what happened on the receive, on the unbuffered channel.

## Final Summary for Channel

Till now we saw the internal working of the channels, what happens when a channel is created?

What happens when we do send or receive on a channel?

Let us summarize.

hchan struct represents the channel, it contains circular ring, buffer and mutex lock. The goroutines, have to acquire the mutex lock to do any channel operation.

When a goroutine gets blocked on send or receive, then they are parked in the send queue or the receive queue.

Go scheduler moves the blocked goroutine out of the OS thread.

Once the channel operation is complete, goroutines are moved back to the local run queue.

This was all about how to channel, send and receive works.

By now, you should have become very comfortable with goroutines and channels, which are the pillars of concurrency in Go.

## Select

Here is our scenario, goroutine G1 has spawned two goroutines, G2 and G3 and has given them a task to do.

![alt text](image-75.png)

Now, the question is, in what order are we going to receive the results from these two goroutines?

Are we going to receive from G2 first and then G3 or G3 first and then G2.

![alt text](image-76.png)

What if the G3 executes faster in some instances and returns the result faster and G2 executes faster in other instances and returns the results faster

So the question is, can we do the operation on the channel, whichever is ready and not worry about the order.

And this is where select comes into play. Select is like a switch statement, each case statement specifies a send or receive on some channel and it has an associated block of statements.

![alt text](image-77.png)

Each case statement is not evaluated sequentially, but all channel operations are considered simultaneously to see if any of them is ready.

And each case has an equal chance of being selected.

Select waits until some case is ready to proceed, if none of the channels are ready, the entire select statement is going to block, until some case is ready for the communication.

When one channel is ready, then it performs the channel operation and executes the associated block of statements.

If multiple channels are ready, then it's going to pick one of them at random.

`Select` is very helpful in implementing timeouts and not blocking communication.

You can specify time out on the channel operation by using select and time after function.

![alt text](image-78.png)

Select will wait until there is a event on the channel or until the timeout is reached.

The time after function will take a time duration as input, and it returns a channel.

And it starts a goroutine in the background, and sends the value on the channel after the specified time duration. In this code snippet, the select will wait for a value to arrive on the channel ch for three seconds.

If it does not arrive, then it's going to get timed out.

you know, channels are blocking, right, you can achieve non-blocking operation with select

by specifying the default case.

![alt text](image-79.png)

If none of the channel operation is ready, then the default case gets executed and the select does not wait for the channel.

It just checks if the operation is ready, if it is, it performs the operation, if not, then the default case gets executed.

So in this code snippet, if some goroutine has already sent a value on the channel ch, then it will read the value.

If there was no goroutine, which has sent a value, then it just executes the default case.

Some scenarios to consider are, the empty select statement will block forever and select on the nil channel will also block forever.
![alt text](image-80.png)

So let us summarize.

So you saw select is like a switch statement.

With each case statement specifying a channel operation.

And the select is going to block until there is any case ready for the channel operation. With select we can implement a timeout and nonblocking communication and select on nil channel will block forever.

### `Select` Exercise:

### Exercise Channel ownership 01-exercise/03-select/01-select

### Timeout Exercise:

### Exercise Channel ownership 01-exercise/03-select/02-select

### Non-blocking Exercise:

### Exercise Channel ownership 01-exercise/03-select/03-select

## sync Package

We have already seen sync wait group, we will look into other utilities in the sync package, like the mutex, condition variables, atomic and pool.

We'll start with mutex (`sync.Mutex`).

Question is when to use channels and when to use mutex?
channels are great for communication between the goroutines.

But what if we have like caches, registries and state, which are big to be sent over the channel and we want the access to these data to be concurrent safe, so that only one goroutine has an access at a time.

So this is where classical synchronization tools like the mutex comes into the picture.

So we use channels to pass data between the goroutines and distribute units of work and communicate asynchronous results, and we use mutex to protect caches, registries and states from concurrent access.

Mutex is used to guard access to the shared resource.

Mutex provides a convention for the developers to follow, anytime a developer wants to access the shared memory, they must first acquire a lock and when they are finished, they must release the lock.

![alt text](image-81.png)

And locks are exclusive, if a goroutine has acquired the lock, then other goroutines will block until the lock is available.

The region between the lock and unlock is called the critical section.

And it is common idiom to call unlock with defer, so that unlock gets executed at the end of the function.

The critical section reflects the bottleneck where only one goroutine can be either be reading or writing to a shared memory, if the goroutine is just reading and not writing to the memory, then we can use read write mutex.

`sync.RWMutex` allows multiple readers access to the critical section, simultaneously, unless the lock is being held by the writer.

![alt text](image-82.png)

The writer gets the exclusive look. And here they defer unlock runs after the return statement has read the value of the balance.

Let us summarize.

Mutex is used to guard access to the shared resource.

It is a developers convention to call lock to access the shared memory and call unlock when done.

And the critical section represents the bottleneck between the goroutines.

### Exercise Mutex:

### Exercise Channel ownership 01-exercise/03-select/03-mutex

`sync.AUtomic` :

Atomic is used to perform low level atomic operation on the memory. It is used by other synchronization utilities.

It is a Lockless operation.

Here in this example, we are using a atomic operation on the counters, we use add method to increment the value of the counter, and this add method can be called by multiple goroutines concurrently and the access to the memory will be concurrent safe. And we use the load method to read the value of the counter in a concurrent safe manner.

`runtime.GOMAXPROCS(4)`: tells go runtime to use 4 cpu cores to run our go routines. so 4 goroutines can be run in parallel

![alt text](image-83.png)

### Exercise Atomic:

### Exercise Channel ownership 01-exercise/04-sync/11-atomic

`sync.Cond` :

Conditional variable is one of the synchronization mechanisms, a conditional variable is basically a container of goroutines that are waiting for a certain condition.

The question is, how can we make a goroutine wait till some event or condition occurs, one way could be to wait in a loop for the condition to become true.

![alt text](image-84.png)

In this code snippet, we have a shared resource, a map that is being shared between the goroutines.

And the consumer goroutine, needs to wait for the shared map to be populated before processing it.

So first we will acquire a lock.

We check for the condition whether the shared map is populated by checking the length of the map.

If it is not populated, then we release the lock, sleep for an arbitrary duration of time and again acquire a lock.

And check for the condition again. This is quite inefficient, right?

What we need is we need some way to make the goroutine suspend while waiting, and some way to signal the suspended goroutine that, that particular event has occurred.

Can we use channels?

We can use channels to block the goroutine on receive and sender goroutine to indicate the occurrence of the event.

But what if there are multiple goroutines waiting on multiple conditions?

That's where conditional variables comes into the picture.

```go
  var c *sync.Cond
```

Conditional variables are of type sync.Cond, we use the constructor method, NewCond() to create a conditional variable, and it takes a sync locker interface as input, which is usually a sync mutex.

![alt text](image-85.png)

And this is what allows the conditional variable to facilitate the coordination between the goroutines in a concurrent. safe way.

sync.Cond package contains three methods.

![alt text](image-86.png)

wait, signal and broadcast.

wait method, suspends the execution of the calling thread, and it automatically releases the lock before suspending the goroutine. Wait does not return unless it is woken up by a broadcast or a signal.

![alt text](image-87.png)

Once it is woken up, it acquires the lock again.

And on resume, the caller should check for the condition again, as it is very much possible that another goroutine could get scheduled between the signal and the resumption of wait and change the state of the condition.

So this is why we check for the condition in a for loop here.

![alt text](image-88.png)

Signal, signal wakes up one goroutine that was waiting on a condition. The signal finds a goroutine that was waiting the longest and notifies that goroutine. And it is allowed, but not required for the caller to hold the lock during this call.

![alt text](image-89.png)

Broadcast, broadcast wakes up all the goroutine that were waiting on the condition, and again, it is allowed, but it is not required for the caller to hold the lock during this call.

Let us look into an example, we have a goroutine G2 which needs to wait for the shared resource to be populated before proceeding with its processing.

![alt text](image-90.png)

We create a conditional variable with the constructor NewCond, passing it the mutex as the input. Here we have our shared resource, this is our goroutine and we take the lock for the entire duration of our processing. We check for the condition, whether they shared resources populated, if not, we make a call to the wait. Wait implicitly releases the lock and suspends our goroutine.

Now the producer goroutine comes along.

![alt text](image-91.png)

It acquires a lock, populates the shared resource, and sends a signal to the consumer goroutine and then it releases the lock.

On receiving the signal, the consumer goroutine is put back to the runnable state and wait acquires the lock again, the wait returns, we check for the condition again, and then we proceed with our processing and we release the lock.

So this is how the wait and signal mechanism works.

If there are multiple goroutines, waiting on a condition, then we use broadcast.

![alt text](image-92.png)

The broadcast will send a signal to all the goroutines that were waiting on the condition. So in this way, we are able to coordinate the execution of the goroutines, when they need to wait on an occurrence of a condition or an event.

Let us summarize, so we saw that conditional variable are used to synchronize the execution of the goroutines, and there are three methods, wait suspends the execution of the goroutine. Signal wakes up one goroutine that was waiting on the condition. Broadcast, wakes up all the goroutines that were waiting on that condition.

### Exercise Cond:

### Exercise Channel ownership 01-exercise/04-sync/21-cond

### Exercise Channel ownership 01-exercise/04-sync/22-cond

`sync.Once`

sync once is used to run one time initialization functions, the Do method accepts the initialization function as its argument. sync once ensures that only one call to Do ever calls the function, that is passed in, even when called from different goroutines.

And this is pretty useful in the creation of a singleton object or calling initialization functions, which multiple goroutines depends on, but we want the initialization function to run only once.

### Exercise Channel ownership 01-exercise/04-sync/31-once

`sync.Pool`

Pool is commonly used to constrain the creation of expensive resources like the database connections, network connections and memory.

We will maintain a pool of fixed number of instances of the resource and those resources from the pool will be reused rather than creating new instances each time whenever the caller requires them.

![alt text](image-94.png)

The caller, calls the get method, whenever it wants access to the resource. And this method will first check, whether there is any available instance within the pool.

If yes, then it returns that instance to the caller.

If not, then a new instance is created which is returned to the caller.

When finished with the usage, the caller, calls the put method, which places the instance back to the pool, which can be reused by other processes.

![alt text](image-95.png)

In this way, you can place a constraint on the creation of expensive resources.

### Exercise Channel ownership 01-exercise/04-sync/41-pool
