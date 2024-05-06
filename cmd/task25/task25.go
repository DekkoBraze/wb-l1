package main

import (
	"time"
)

// Функция Sleep на основе таймера
func SleepSeconds(duration int) {
	seconds := time.Second * time.Duration(duration)
	timer := time.NewTimer(seconds)
	<-timer.C
}

func main() {
	SleepSeconds(5)
}
