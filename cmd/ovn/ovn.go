package main

import (
	"fmt"
	"github.com/sgkind/ovn-viewer/pkg/nbdb"

	goovn "github.com/ebay/go-ovn"
)

const (
	ovsRundir = "/var/run/openvswitch"
	ovnnbSocket = "ovnnb_db.sock"
)

var ovn goovn.Client

func main() {
	var err error
	nbController := nbdb.NewController()
	ovn, err = goovn.NewClient(
		&goovn.Config{
			Addr: "tcp:192.168.200.101:6641",
			SignalCB: nbController,
		})
	if err != nil {
		panic(err)
	}

	switches := nbController.GetLogicalSwitch()
	for _, ls := range switches {
		fmt.Printf("%s\n", ls.Name)
	}

	//switches, _ := ovn.LSList()

	//for _, ls := range switches {
	//	fmt.Printf("%v\n", ls)
	//}
}
