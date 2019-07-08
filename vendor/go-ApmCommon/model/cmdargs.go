package model

import (
	"flag"
	"fmt"
	"sync"
)

var ca cmdargs

type cmdargs struct {
	Phase string
}

func (ca cmdargs) String() string {
	return fmt.Sprintf("phase : %s", ca.Phase)
}

// init is a function to parse flag
func init() {
	p := flag.String("phase", "local", "input phase e.g)local, dv")
	flag.Parse()
	ca = cmdargs{
		Phase: *p,
	}
}

// Singleton
var insCmdargs *cmdargs
var onceCmdargs sync.Once

func GetCmdargs() *cmdargs {
	onceCmdargs.Do(func() {
		insCmdargs = &ca
	})
	return insCmdargs
}
