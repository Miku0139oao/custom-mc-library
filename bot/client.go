package bot

import (
	"github.com/rain931215/custom-mc-library/net"
)

// Client is used to access Minecraft server
type Client struct {
	conn *net.Conn
	Auth
}
