package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	FinalWord = "GO!"
	CountStart = 3
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct{
	duration time.Duration
}

func (o ConfigurableSleeper) Sleep(){
	time.Sleep(o.duration)
}

func Countdown(out io.Writer, sleeper Sleeper){
	for i := CountStart; i  >0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out,i)
	}

	sleeper.Sleep()
	fmt.Fprint(out,FinalWord)
}

func main(){
	sleeper:=&ConfigurableSleeper{1*time.Second}
	Countdown(os.Stdout,sleeper)
}