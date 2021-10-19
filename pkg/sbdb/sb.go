package sbdb

import (

	ovn "github.com/ebay/go-ovn"
)

type Controller struct {
	client              ovn.Client

	chassisMap          map[string]*ovn.Chassis
	chassisPrivateMap   map[string]*ovn.ChassisPrivate
	encapMap            map[string]map[string]*ovn.Encap
}

var _ ovn.OVNSignal = &Controller{}
