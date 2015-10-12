// +build linux

package drivers

import (
	"errors"

	log "github.com/Sirupsen/logrus"
	"github.com/contiv/netplugin/mgmtfn/dockplugin/libnetClient"
)

func createDockNet(networkName string) error {
	api := libnetClient.NewRemoteAPI("")

	// Check if the network already exists
	nw, err := api.NetworkByName(networkName)
	if err == nil && nw.Type() == driverName {
		return nil
	} else if err == nil && nw.Type() != driverName {
		log.Errorf("Network name %s used by another driver %s", networkName, nw.Type())
		return errors.New("Network name used by another driver")
	}

	// Create network
	_, err = api.NewNetwork(driverName, networkName)
	if err != nil {
		log.Errorf("Error creating network %s. Err: %v", networkName, err)
		// FIXME: Ignore errors till we fully move to docker 1.9
		return nil
	}

	return nil
}
