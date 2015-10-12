// +build !windows !darwin

// Package netlinkutils implements netlink functionality for the drivers.
package netlinkutils

import (
	log "github.com/Sirupsen/logrus"
	"github.com/contiv/netplugin/netutils"
	"github.com/vishvananda/netlink"
)

// SetInterfaceMac wraps netutils.SetInterfaceMac
func SetInterfaceMac(intfName string, macaddr string) error {
	return netutils.SetInterfaceMac(intfName, macaddr)
}

// CreateVethPair creates veth interface pairs with specified name
func CreateVethPair(name1, name2 string) error {
	log.Infof("Creating Veth pairs with name: %s, %s", name1, name2)

	// Veth pair params
	veth := &netlink.Veth{
		LinkAttrs: netlink.LinkAttrs{
			Name:   name1,
			TxQLen: 0,
		},
		PeerName: name2,
	}

	// Create the veth pair
	if err := netlink.LinkAdd(veth); err != nil {
		log.Errorf("error creating veth pair: %v", err)
		return err
	}

	return nil
}

// DeleteVethPair deletes veth interface pairs
func DeleteVethPair(name1, name2 string) error {
	log.Infof("Deleting Veth pairs with name: %s, %s", name1, name2)

	// Veth pair params
	veth := &netlink.Veth{
		LinkAttrs: netlink.LinkAttrs{
			Name:   name1,
			TxQLen: 0,
		},
		PeerName: name2,
	}

	// Create the veth pair
	if err := netlink.LinkDel(veth); err != nil {
		log.Errorf("error deleting veth pair: %v", err)
		return err
	}

	return nil
}

// SetLinkUp sets the link up
func SetLinkUp(name string) error {
	iface, err := netlink.LinkByName(name)
	if err != nil {
		return err
	}
	return netlink.LinkSetUp(iface)
}

// SetLinkMtu Sets the link MTU
func SetLinkMtu(name string, mtu int) error {
	iface, err := netlink.LinkByName(name)
	if err != nil {
		return err
	}
	return netlink.LinkSetMTU(iface, mtu)
}
