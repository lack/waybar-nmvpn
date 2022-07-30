package main

import (
	"fmt"
	"time"

	"github.com/lack/waybar-nmvpn/pkg/nmvpn"

	waybar "github.com/lack/gowaybarplug"
)

func loop(interval time.Duration) {
	wb := waybar.NewUpdater()

	for true {
		status := waybar.Status{
			Text: "VPN",
		}
		vpns, err := nmvpn.GetVPNs()
		if err == nil && len(vpns) > 0 {
			active := false
			for _, v := range vpns {
				if v.Active {
					active = true
					break
				}
			}
			if active {
				status.Alt = "connected"
				status.Class = []string{"connected"}
			} else {
				status.Alt = "disconnected"
				status.Class = []string{"disconnected"}
			}
			status.Tooltip = ""
			for _, v := range vpns {
				if status.Tooltip != "" {
					status.Tooltip += "\n"
				}
				state := "down"
				if v.Active {
					state = "up"
				}
				status.Tooltip += fmt.Sprintf("%s: %s", v.Name, state)
			}
		} else if err != nil {
			status.Alt = "error"
			status.Class = []string{"error"}
			status.Tooltip = fmt.Sprintf("Error fetching VPN information: %v", err)
		} else {
			status.Alt = "none"
			status.Class = []string{"unconfigured"}
			status.Tooltip = "No VPN connections configured"
		}
		wb.Status <- &status
		time.Sleep(interval)
	}
}

func main() {
	loop(5 * time.Second)
}
