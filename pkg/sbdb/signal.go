package sbdb

import (
	ovn "github.com/ebay/go-ovn"
	"github.com/golang/glog"
)

func (sc *Controller) OnLogicalSwitchCreate(ls *ovn.LogicalSwitch) {
    glog.Errorf("got notification on creating logical switch-%s from south bound db", ls.Name)
}

func (sc *Controller) OnLogicalSwitchDelete(ls *ovn.LogicalSwitch) {
    glog.Errorf("got notification on deleting logical switch-%s from south bound db", ls.Name)
}

func (sc *Controller) OnLogicalPortCreate(lp *ovn.LogicalSwitchPort) {
    glog.Errorf("got notification on creating logical switch port-%s from south bound db", lp.Name)
}

func (sc *Controller) OnLogicalPortDelete(lp *ovn.LogicalSwitchPort) {
    glog.Errorf("got notification on deleting logical switch port-%s from south bound db", lp.Name)
}

func (sc *Controller) OnLogicalRouterCreate(lr *ovn.LogicalRouter) {
    glog.Errorf("got notification on creating logical router-%s from south bound db", lr.Name)
}

func (sc *Controller) OnLogicalRouterDelete(lr *ovn.LogicalRouter) {
    glog.Errorf("got notification on deleting logical router-%s from south bound db", lr.Name)
}

func (sc *Controller) OnLogicalRouterPortCreate(lrp *ovn.LogicalRouterPort) {
	glog.Errorf("got notification on creating logical router port-%s from south bound db",
		lrp.Name)
}

func (sc *Controller) OnLogicalRouterPortDelete(lrp *ovn.LogicalRouterPort) {
	glog.Errorf("got notification on deleting logical router port-%s from south bound db",
		lrp.Name)
}

func (sc *Controller) OnLogicalRouterStaticRouteCreate(lrsr *ovn.LogicalRouterStaticRoute) {
	glog.Errorf("got notification on creating logical router static route(%s) from south bound db",
		lrsr)
}

func (sc *Controller) OnLogicalRouterStaticRouteDelete(lrsr *ovn.LogicalRouterStaticRoute) {
	glog.Errorf("got notification on deleting logical router static route(%s) from south bound db",
		lrsr)
}

func (sc *Controller) OnACLCreate(acl *ovn.ACL) {
    glog.Errorf("got notification on creating acl-%s from south bound db", acl.Name)
}

func (sc *Controller) OnACLDelete(acl *ovn.ACL) {
    glog.Errorf("got notification on deleting acl-%s from south bound db", acl.Name)
}

func (sc *Controller) OnDHCPOptionsCreate(dhcp *ovn.DHCPOptions) {
    glog.Errorf("got notification on creating dhcp(%v) from south bound db", dhcp)
}

func (sc *Controller) OnDHCPOptionsDelete(dhcp *ovn.DHCPOptions) {
	glog.Errorf("got notification on deleting dhcp(%v) from south bound db", dhcp)
}

func (sc *Controller) OnQoSCreate(qos *ovn.QoS) {
    glog.Errorf("got notification on creating qos(%v) from south bound db", qos)
}

func (sc *Controller) OnQoSDelete(qos *ovn.QoS) {
	glog.Errorf("got notification on deleting qos(%v) from south bound db", qos)
}

func (sc *Controller) OnLoadBalancerCreate(lb *ovn.LoadBalancer) {
    glog.Errorf("got notification on creating load balancer-%s from south bound db", lb.Name)
}

func (sc *Controller) OnLoadBalancerDelete(lb *ovn.LoadBalancer) {
    glog.Errorf("got notification on deleting load balancer-%s from south bound db", lb.Name)
}

func (sc *Controller) OnMeterCreate(meter *ovn.Meter) {
    glog.Errorf("got notification on creating meter-%s from south bound db", meter.Name)
}

func (sc *Controller) OnMeterDelete(meter *ovn.Meter) {
    glog.Errorf("got notification on deleting meter-%s from south bound db", meter.Name)
}

func (sc *Controller) OnMeterBandCreate(band *ovn.MeterBand) {
    glog.Errorf("got notification on creating meter band(%v) from south bound db", band)
}

func (sc *Controller) OnMeterBandDelete(band *ovn.MeterBand) {
    glog.Errorf("got notification on deleting meter band(%v) from south bound db", band)
}

func (sc *Controller) OnChassisCreate(ch *ovn.Chassis) {
    glog.V(2).Infof("chassis-%s is created", ch.Name)
	sc.chassisMap[ch.UUID] = ch
}

func (sc *Controller) OnChassisDelete(ch *ovn.Chassis) {
    glog.V(2).Infof("chassis-%s is deleted", ch.Name)
	delete(sc.chassisMap, ch.UUID)
}

func (sc *Controller) OnEncapCreate(encap *ovn.Encap) {
	glog.V(2).Infof("encap(%s) is created", encap)
	var ecs map[string]*ovn.Encap
    glog.V(2).Infof("encap(%v) is created", encap)
	if ecs, ok := sc.encapMap[encap.ChassisName]; !ok {
		ecs = make(map[string]*ovn.Encap)
		sc.encapMap[encap.ChassisName] = ecs
	}

	ecs[encap.UUID] = encap
}

func (sc *Controller) OnEncapDelete(encap *ovn.Encap) {
    glog.V(2).Infof("encap(%v) is deleting", encap)
	ecs, ok := sc.encapMap[encap.ChassisName]
	if !ok {
		return
	}

	delete(ecs, encap.UUID)
}
