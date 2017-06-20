package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var status = make(chan string)

func main() {
	wg.Add(2)
	go salvaNoBd()
	go salvaNoRedis()
	wg.Wait()
}

func salvaNoBd() {
	status <- "Func 1"
	time.Sleep(3 * time.Second)
	fmt.Println("Acabei de salvar no BD")
	wg.Done()
}

func salvaNoRedis() {
	msg := <-status
	fmt.Println("Qual a msg para o Redis? ", msg)
	time.Sleep(1 * time.Second)
	fmt.Println("Acabei de salvar no Redis")
	close(status)
	wg.Done()
}
