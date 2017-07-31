package main

/*
#include <sys/types.h>
#include <pwd.h>
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type Passwd struct {
	Uid, Gid   uint32
	Dir, Shell string
}

func Getpwnam(name string) *Passwd {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cpw := C.getpwnam(cname)
	return &Passwd{
		Uid: uint32(cpw.pw_uid), Gid: uint32(cpw.pw_uid),
		Dir: C.GoString(cpw.pw_dir), Shell: C.GoString(cpw.pw_shell)}
}

func main() {
	fmt.Print(Getpwnam("brrichin"))
}
