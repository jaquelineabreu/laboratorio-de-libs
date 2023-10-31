package main

import (
	"fmt"

	"gopkg.in/robfig/cron.v2"
)

func hello(name string) {
	message := fmt.Sprintf("Hi, %v", name)
	fmt.Println(message)
}

func runCronJobs() {
	s := cron.New()

	s.AddFunc("@every 1s", func() {
		hello("Cron v2")
	})

	s.Start()
}

func main() {
	runCronJobs()
	fmt.Scanln()

}
