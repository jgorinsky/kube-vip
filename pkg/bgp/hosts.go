package bgp

import (
	"context"
	"fmt"
	"net/netip"

	api "github.com/osrg/gobgp/v3/api"
	log "github.com/sirupsen/logrus"
)

// AddHost will update peers of a host
func (b *Server) AddHost(addr string) (err error) {
	ip, err := netip.ParsePrefix(addr)
	if err != nil {
		return err
	}

	log.Debugf("JG Parsed Prefix is [%s]", ip)

	p := b.getPath(ip)
	if p == nil {
		return fmt.Errorf("failed to get path for %v", ip)
	}
	log.Debugf("JG Path is [%s]", p)

	_, err = b.s.AddPath(context.Background(), &api.AddPathRequest{
		Path: p,
	})

	if err != nil {
		return err
	}

	return
}

// DelHost will inform peers to remove a host
func (b *Server) DelHost(addr string) (err error) {
	ip, err := netip.ParsePrefix(addr)
	if err != nil {
		return err
	}
	p := b.getPath(ip)
	if p == nil {
		return
	}

	return b.s.DeletePath(context.Background(), &api.DeletePathRequest{
		Path: p,
	})
}
