package main

import (
	"./network/network/bcast"
	"./network/network/localip"
	"flag"
	"fmt"
	"os"
	"time"
	"./timer"
	"log"
	"os/exec"
)

type backupInfo struct {
	Message string
	It int
}

func spawnBackup(){
	cmd:=exec:Command("gnome-terminal", "--window", "-x", "go", "run", "Primary.go")
}

func main(){
	count := 0
	timer := timer.timerstruct{}
	timer.init_timer(3000, &timer)

	timeout := make(chan bool)
	Rx := make(chan BackupInfo)

	go bcast.Receiver(16569, Rx)
	go timer.poll_timer(timeout, &timer);

	fmt.Println("\n backup receiving data from primary\n")
 backup := true
	for backup{
		select {
		case a := <-Rx:
			count = a.It
			timer.reset(&timer)
		case <- timeout
			fmt.Println("timed out, primary assumed dead. You have been gnomed")
		}
	}
	spawnBackup()

	Tx := make(chan backupInfo)
	go bcast.Transmitter(16569, Tx)

	go func runBackup(){
		BackupInfo := BackupInfo{"From program " + id, count}
		for {
			Tx <- BackupInfo
			BackupMsg.It++
			time.Sleep(1 * time.Second)
			fmt.Println("Iteration variable:", BackupMsg.It)
		}
	}()
	}
}
