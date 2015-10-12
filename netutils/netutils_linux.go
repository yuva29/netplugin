// +build linux

// Package netutils contains utility functions for manipulating linux networks.
package netutils

import (
	"net"

	"github.com/contiv/netplugin/core"

	"github.com/vishvananda/netlink"
)

// GetInterfaceIP obtains the ip addr of a local interface on the host.
func GetInterfaceIP(linkName string) (string, error) {
	var addrs []netlink.Addr
	localIPAddr := ""

	link, err := netlink.LinkByName(linkName)
	if err != nil {
		return "", err
	}
	addrs, err = netlink.AddrList(link, netlink.FAMILY_V4)
	if err != nil {
		return "", err
	}
	if len(addrs) > 0 {
		localIPAddr = addrs[0].IP.String()
	}

	err = core.Errorf("local ip not found")
	if localIPAddr != "" {
		err = nil
	}

	return localIPAddr, err
}

// SetInterfaceMac : Set mac address of an interface
func SetInterfaceMac(name string, macaddr string) error {
	iface, err := netlink.LinkByName(name)
	if err != nil {
		return err
	}
	hwaddr, err := net.ParseMAC(macaddr)
	if err != nil {
		return err
	}
	return netlink.LinkSetHardwareAddr(iface, hwaddr)
}
