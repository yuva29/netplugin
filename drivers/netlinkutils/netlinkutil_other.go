// +build !linux

// Package netlinkutils implements netlink functionality for the drivers.
package netlinkutils

// SetInterfaceMac wraps netutils.SetInterfaceMac
func SetInterfaceMac(intfName string, macaddr string) error {
	panic("Setting interface properties is not supported on that platform")
}

// CreateVethPair creates veth interface pairs with specified name
func CreateVethPair(name1, name2 string) error {
	panic("Creating veth devices not supported on this platform")
}

// DeleteVethPair deletes veth interface pairs
func DeleteVethPair(name1, name2 string) error {
	panic("Creating veth devices not supported on this platform")
}

// SetLinkUp sets the link up
func SetLinkUp(name string) error {
	panic("Creating veth devices not supported on this platform")
}

// SetLinkMtu Sets the link MTU
func SetLinkMtu(name string, mtu int) error {
	panic("Manipulating veth devices not supported on this platform")
}
