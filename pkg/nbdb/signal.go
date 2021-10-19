package nbdb

import (
	ovn "github.com/ebay/go-ovn"
	"github.com/golang/glog"
)

func (nc *Controller) OnLogicalSwitchCreate(ls *ovn.LogicalSwitch) {
    glog.V(2).Infof("logical switch-%s is created\n", ls.Name)
	nc.switches[ls.UUID] = ls
}

func (nc *Controller) OnLogicalSwitchDelete(ls *ovn.LogicalSwitch) {
    glog.V(2).Infof("logical switch-%s is deleted\n", ls.Name)
	delete(nc.switches, ls.UUID)
	// TODO: delete logical switch port
}

func (nc *Controller) OnLogicalPortCreate(lp *ovn.LogicalSwitchPort) {
    glog.V(2).Infof("logical switch port-%s is created\n", lp.Name)
	nc.lsPorts[lp.UUID] = lp
}

func (nc *Controller) OnLogicalPortDelete(lp *ovn.LogicalSwitchPort) {
    glog.V(2).Infof("logical switch port-%s is deleted\n", lp.Name)
	delete(nc.lsPorts, lp.UUID)
}

func (nc *Controller) OnLogicalRouterCreate(lr *ovn.LogicalRouter) {
	glog.V(2).Infof("logical router-%s is created\n", lr.Name)
	nc.routers[lr.UUID] = lr
}

func (nc *Controller) OnLogicalRouterDelete(lr *ovn.LogicalRouter) {
	glog.V(2).Infof("logical router-%s is deleted\n", lr.Name)
	delete(nc.routers, lr.Name)
	// TODO: delete logical router port
}

func (nc *Controller) OnLogicalRouterPortCreate(lrp *ovn.LogicalRouterPort) {
    glog.V(2).Infof("logical router port-%s is created\n", lrp.Name)
	nc.lrPorts[lrp.UUID] = lrp
}

func (nc *Controller) OnLogicalRouterPortDelete(lrp *ovn.LogicalRouterPort) {
    glog.V(2).Infof("logical router port-%s is deleted\n", lrp.Name)
	delete(nc.lrPorts, lrp.UUID)
}

func (nc *Controller) OnLogicalRouterStaticRouteCreate(lrsr *ovn.LogicalRouterStaticRoute) {
    glog.V(2).Infof("logical router static route(%v) is created\n", lrsr)
	nc.lrRoutes[lrsr.UUID] = lrsr
}

func (nc *Controller) OnLogicalRouterStaticRouteDelete(lrsr *ovn.LogicalRouterStaticRoute) {
    glog.V(2).Infof("logical router static route(%v) is deleted\n", lrsr)
	delete(nc.lrRoutes, lrsr.UUID)
}

func (nc *Controller) OnACLCreate(acl *ovn.ACL) {
    //glog.V(2).Infof("acl-%s is created\n", acl.Name)
	//nc.acls[acl.UUID] = acl
}

func (nc *Controller) OnACLDelete(acl *ovn.ACL) {
    //glog.V(2).Infof("acl-%s is deleted\n", acl.Name)
	//delete(nc.acls, acl.UUID)
}

func (nc *Controller) OnDHCPOptionsCreate(dhcp *ovn.DHCPOptions) {
    glog.V(2).Infof("dhcp(%v) is created\n", dhcp)
	nc.dhcpOptions[dhcp.UUID] = dhcp
}

func (nc *Controller) OnDHCPOptionsDelete(dhcp *ovn.DHCPOptions) {
    glog.V(2).Infof("dhcp(%v) is deleted\n", dhcp)
	delete(nc.dhcpOptions, dhcp.UUID)
}

func (nc *Controller) OnQoSCreate(qos *ovn.QoS) {
    glog.V(2).Infof("qos(%v) is created\n", qos)
	nc.qosMap[qos.UUID] = qos
}

func (nc *Controller) OnQoSDelete(qos *ovn.QoS) {
    glog.V(2).Infof("qos(%v) is deleted\n", qos)
	delete(nc.qosMap, qos.UUID)
}

func (nc *Controller) OnLoadBalancerCreate(ls *ovn.LoadBalancer) {
    glog.V(2).Infof("load balancer-%s is created\n", ls.Name)
	nc.loadBalancers[ls.UUID] = ls
}

func (nc *Controller) OnLoadBalancerDelete(ls *ovn.LoadBalancer) {
    glog.V(2).Infof("load balancer-%s is deleted\n", ls.Name)
	delete(nc.loadBalancers, ls.UUID)
}

func (nc *Controller) OnMeterCreate(meter *ovn.Meter) {
    glog.V(2).Infof("meter-%s is created\n", meter.Name)
	nc.meters[meter.UUID] = meter
}

func (nc *Controller) OnMeterDelete(meter *ovn.Meter) {
    glog.V(2).Infof("meter-%s is deleted\n", meter.Name)
	delete(nc.meters, meter.UUID)
}

func (nc *Controller) OnMeterBandCreate(band *ovn.MeterBand) {
    glog.V(2).Infof("meter band(%v) is created\n", band)
	nc.meterBands[band.UUID] = band
}

func (nc *Controller) OnMeterBandDelete(band *ovn.MeterBand) {
    glog.V(2).Infof("meter band(%v) is deleted\n", band)
	delete(nc.meterBands, band.UUID)
}

func (nc *Controller) OnChassisCreate(ch *ovn.Chassis) {
    glog.Errorf("got notification on creating chassis-%s from north bound db", ch.Name)
}

func (nc *Controller) OnChassisDelete(ch *ovn.Chassis) {
    glog.Errorf("got notification on deleting chassis-%s from north bound db", ch.Name)
}

func (nc *Controller) OnEncapCreate(encap *ovn.Encap) {
    glog.Errorf("got notification on creating encap(%v) from north bound db", encap)
}

func (nc *Controller) OnEncapDelete(encap *ovn.Encap) {
    glog.Errorf("got notification on deleting encap(%v) from north bound db", encap)
}