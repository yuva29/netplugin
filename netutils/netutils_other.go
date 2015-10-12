// +build !linux

// Package netutils contains utility functions for manipulating linux networks.
package netutils

func GetInterfaceIP(linkName string) (string, error) {
	panic("not implemented")
}

func SetInterfaceMac(name string, macaddr string) error {
	panic("not implemented")
}
