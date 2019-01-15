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

type passwd struct {
	UID, Gid   uint32
	Dir, Shell string
}

func getpwnam(name string) *passwd {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cpw := C.getpwnam(cname)
	return &passwd{
		UID: uint32(cpw.pw_uid), Gid: uint32(cpw.pw_uid),
		Dir: C.GoString(cpw.pw_dir), Shell: C.GoString(cpw.pw_shell)}
}

func main() {
	fmt.Print(getpwnam("brichins"))
}
