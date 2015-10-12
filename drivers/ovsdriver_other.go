// +build !linux

package drivers

func createDockNet(networkName string) error {
	panic("Unsupported on this platform")
}
