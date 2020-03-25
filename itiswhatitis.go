package main

import (
	"log"
	"math/rand"
	"os/exec"
	"sync"
	"time"
)

func itis(delay int) {
	<-time.After(time.Duration(delay) * time.Millisecond)
	s := "eed ees whad eed ees"
	cmd := exec.Command("say", s)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())
	itis(0)
	for i := 0; i < 10; i++ {
		v := rand.Intn(100)
		wg.Add(1)
		go func() {
			defer wg.Done()
			itis(v)
		}()
	}
	wg.Wait()
}
