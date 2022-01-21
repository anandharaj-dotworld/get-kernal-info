package kernal

import (
	"fmt"
	"syscall"
)

func int8ToStr(arr []int8) string {
	b := make([]byte, 0, len(arr))
	for _, v := range arr {
		if v == 0x00 {
			break
		}
		b = append(b, byte(v))
	}
	return string(b)
}

func Kernal() {

	var uname syscall.Utsname
	if err := syscall.Uname(&uname); err == nil {
		fmt.Println("sysname: ", int8ToStr(uname.Sysname[:]),
			"release: ", int8ToStr(uname.Release[:]),
			"version: ", int8ToStr(uname.Version[:]),
			"Arch: ", int8ToStr(uname.Machine[:]),
			"Nname: ", int8ToStr(uname.Nodename[:]))
	}
}
