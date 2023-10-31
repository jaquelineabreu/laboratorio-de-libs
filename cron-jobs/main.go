package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/robfig/cron"
)

func main() {
	simpleCron()
	listen()

}

// Schedule a method to run periodically
func simpleCron() {
	fmt.Println(time.Now().String() + " Start App - Version SimpleCron")
	c := cron.New()
	c.AddFunc("@every 1m10s", task)
	go c.Start()
}

func task() {
	fmt.Println(time.Now().String() + " - Start Task")
	time.Sleep(1 * time.Second)
	fmt.Println(time.Now().String() + " - End Task")
}

// channel that will be notified if any system interruption signal occurs
func listen() {
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
	fmt.Println(time.Now().String() + " - closed")
}
