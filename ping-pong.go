package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Ball struct {
	hits int
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	table := make(chan *Ball)

	// Two players on the table.
	go player("ping", table)
	go player("pong", table)

	// Toss the ball onto the table; game on.
	// If we don't toss it here, we'll deadlock.
	table <- new(Ball)

	// Players can play only for a total of 1 second.
	time.Sleep(1 * time.Second)

	// Grab the ball off the table; game ends.
	<-table

	panic("show me the stacks")
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		table <- ball
	}
}
