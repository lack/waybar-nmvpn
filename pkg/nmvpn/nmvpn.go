package nmvpn

import (
	gonm "github.com/Wifx/gonetworkmanager"
)

type Vpn struct {
	Name   string
	Uuid   string
	Active bool
}

func GetVPNs() ([]Vpn, error) {
	settings, err := gonm.NewSettings()
	if err != nil {
		return nil, err
	}

	connections, err := settings.ListConnections()
	if err != nil {
		return nil, err
	}

	vpns := make([]Vpn, 0, len(connections))
	for _, c := range connections {
		s, err := c.GetSettings()
		if err != nil {
			return nil, err
		}
		ctype := s["connection"]["type"].(string)
		if ctype == "vpn" {
			vpn := Vpn{
				Name: s["connection"]["id"].(string),
				Uuid: s["connection"]["uuid"].(string),
			}
			vpns = append(vpns, vpn)
		}
	}
	if len(vpns) > 0 {
		nm, err := gonm.NewNetworkManager()
		if err != nil {
			return nil, err
		}

		actives, err := nm.GetPropertyActiveConnections()
		for _, a := range actives {
			uuid, err := a.GetPropertyUUID()
			if err != nil {
				return nil, err
			}
			for i, v := range vpns {
				if uuid == v.Uuid {
					vpns[i].Active = true
				}
			}
		}
	}
	return vpns, nil
}
