// +build !linux

package docker

func setIfNs(ifname string, pid int) error {
	panic("Not supported on this platform")
}
