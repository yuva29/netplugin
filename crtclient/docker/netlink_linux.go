// +build linux

package docker

import (
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/vishvananda/netlink"
)

func setIfNs(ifname string, pid int) error {
	link, err := netlink.LinkByName(ifname)
	if err != nil {
		if !strings.Contains(err.Error(), "Link not found") {
			log.Errorf("unable to find link %q. Error: %q", ifname, err)
			return err
		}
		// try once more as sometimes (somehow) link creation is taking
		// sometime, causing link not found error
		time.Sleep(1 * time.Second)
		link, err = netlink.LinkByName(ifname)
		if err != nil {
			log.Errorf("unable to find link %q. Error %q", ifname, err)
			return err
		}
	}

	err = netlink.LinkSetNsPid(link, pid)
	if err != nil {
		log.Errorf("unable to move interface '%s' to pid %d. Error: %s",
			ifname, pid, err)
	}

	return err
}
