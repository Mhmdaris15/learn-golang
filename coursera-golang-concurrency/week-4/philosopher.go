package main

import (
	"fmt"
	"sync"
)

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	number         int
	leftChopstick  *Chopstick
	rightChopstick *Chopstick
}

var host = make(chan bool, 2)

func (p Philosopher) Eat(wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		host <- true // Request permission from the host to eat
		p.pickUpChopsticks()
		p.startEating()
		p.finishEating()
		p.putDownChopsticks()
		<-host // Release permission to eat
	}
	wg.Done()
}

func (p Philosopher) pickUpChopsticks() {
	p.leftChopstick.Lock()
	p.rightChopstick.Lock()
}

func (p Philosopher) putDownChopsticks() {
	p.leftChopstick.Unlock()
	p.rightChopstick.Unlock()
}

func (p Philosopher) startEating() {
	fmt.Printf("Philosopher %d starting to eat\n", p.number)
}

func (p Philosopher) finishEating() {
	fmt.Printf("Philosopher %d finishing eating\n", p.number)
}

func main() {
	chopsticks := make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = new(Chopstick)
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			number:         i + 1,
			leftChopstick:  chopsticks[i],
			rightChopstick: chopsticks[(i+1)%5],
		}
	}

	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go philosophers[i].Eat(&wg)
	}

	wg.Wait()
}


/*
In this implementation, each philosopher is represented by a goroutine. The Philosopher struct contains the philosopher's number and references to the left and right chopsticks. The Chopstick struct is implemented using a mutex to ensure exclusive access.The host channel is used to control concurrency and allow a maximum of two philosophers to eat concurrently. The host <- true statement requests permission to eat, and <-host releases the permission. The Eat method represents the main routine for each philosopher. It loops three times, representing the philosopher eating three times. The philosopher requests permission from the host, picks up the chopsticks, starts eating, finishes eating, puts down the chopsticks, and releases the permission. The main function initializes the chopsticks and philosophers, and then starts each philosopher goroutine. It waits for all goroutines to finish using the sync.WaitGroup. When running the program, you will see the philosophers starting and finishing eating in a non-deterministic order based on the available concurrency permitted by the host. Please note that due to the concurrent nature of the problem, the output may vary on each execution.
*/