# waybar-nmvpn

Waybar plugin to display NetworkManager VPN status

## Installation

```
go install github.com/lack/waybar-nmvpn@latest
```

## Configuration

In `$XDG_CONFIG_HOME/waybar/config`
```json
{
    // ... other waybar configuration
    "custom/nmvpn": {
        "format": "{} {icon}",
        "return-type": "json",
        "exec": "$GOPATH/bin/waybar-nmvpn",
        "format-icons": {
            "connected": "",
            "disconnected": "",
            "none": "",
            "error": "⚠"
        }
    }
}
```

In `$XDG_CONFIG_HOME/waybar/style.css`
```css
#custom-nmvpn.connected {
    background-color: rgba(0x29, 0x80, 0xb9, 0.8);
    box-shadow: inset 0 -3px rgba(0x29, 0x80, 0xb9, 1.0);
}

#custom-nmvpn.disconnected {
    background-color: rgba(0xf5, 0x3c, 0x3c, 0.8);
    box-shadow: inset 0 -3px rgba(0xf5, 0x3c, 0x3c, 1.0);
}

#custom-nmvpn.error {
    background-color: rgba(0xeb, 0x4d, 0x4b, 0.8);
    box-shadow: inset 0 -3px rgba(0xeb, 0x4d, 0x4b, 1.0);
}
```
