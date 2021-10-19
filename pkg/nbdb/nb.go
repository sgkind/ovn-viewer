package nbdb

import (

	ovn "github.com/ebay/go-ovn"
)

type Controller struct {
    client     ovn.Client

    switches   map[string]*ovn.LogicalSwitch
	lsPorts    map[string]*ovn.LogicalSwitchPort

	routers    map[string]*ovn.LogicalRouter
	lrPorts    map[string]*ovn.LogicalRouterPort
	lrPolicies map[string]*ovn.LogicalRouterPolicy
	lrRoutes   map[string]*ovn.LogicalRouterStaticRoute

    loadBalancers  map[string]*ovn.LoadBalancer
	natMap         map[string]*ovn.NAT

    meters         map[string]*ovn.Meter
	meterBands     map[string]*ovn.MeterBand
	acls           map[string]*ovn.ACL
	dhcpOptions    map[string]*ovn.DHCPOptions

    portGroups     map[string]*ovn.PortGroup
	qosMap         map[string]*ovn.QoS

	gatewayChassis map[string]*ovn.GatewayChassis
}

var _ ovn.OVNSignal = &Controller{}

func NewController() *Controller {
	return &Controller {
		switches: make(map[string]*ovn.LogicalSwitch),
		lsPorts: make(map[string]*ovn.LogicalSwitchPort),
		routers: make(map[string]*ovn.LogicalRouter),
		lrPorts: make(map[string]*ovn.LogicalRouterPort),
		lrPolicies: make(map[string]*ovn.LogicalRouterPolicy),
		lrRoutes: make(map[string]*ovn.LogicalRouterStaticRoute),
		loadBalancers: make(map[string]*ovn.LoadBalancer),
		natMap: make(map[string]*ovn.NAT),
		meters: make(map[string]*ovn.Meter),
		meterBands: make(map[string]*ovn.MeterBand),
		acls: make(map[string]*ovn.ACL),
		dhcpOptions: make(map[string]*ovn.DHCPOptions),
		portGroups: make(map[string]*ovn.PortGroup),
		qosMap: make(map[string]*ovn.QoS),
		gatewayChassis: make(map[string]*ovn.GatewayChassis),
	}
}

func (nc *Controller) setClient(nbClient ovn.Client) {
	nc.client = nbClient
}

func (nc *Controller) GetLogicalSwitch() []*ovn.LogicalSwitch {
	var switches []*ovn.LogicalSwitch
    for _, ls := range nc.switches {
		switches = append(switches, ls)
	}

	return switches
}