package main

import (
	"fmt"
	"image"
)

var balance int

func Deposit(amount int) { balance = balance + amount }
func Balance() int       { return balance }

func main() {
	// Alice:
	go func() {
		Deposit(200)                // A1
		fmt.Println("=", Balance()) // A2
	}()

	// Bob:
	go Deposit(100) // B


	var deposits = make(chan int) // send amount to deposit
	var balances = make(chan int) // receive balance

	func Deposit(amount int) { deposits <- amount }
	func Balance() int       { return <-balances }

	func teller() {
		var balance int // balance is confined to teller goroutine
		for {
			select {
			case amount := <-deposits:
				balance += amount
			case balances <- balance:
			}
		}
	}

	func init() {
		go teller() // start the monitor goroutine
	}


}
