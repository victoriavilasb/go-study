package custom_net

import "internal/bytealg"

const (
	IPv4len = 4
	IPv6len = 6
)

type IP []byte

type IPMask []byte

type IPNet struct {
	IP IP
	Mask IPMask
}

func IPv4(a, b, c, d byte) IP {

}