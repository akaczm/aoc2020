// Primitive exercise in structures using IPv4 as an example

package main

import (
	"fmt"
	"math"
	"math/bits"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type IPv4 struct {
	addr uint32
	mask uint32
}

func NewIP(address string) (IPv4, error) {
	split := strings.Split(address, "/")
	addr := split[0]
	ipv4, err := decodeAddress(addr)
	if err != nil {
		return IPv4{}, err
	}
	mask := split[1]
	bitmask, err := decodeMask(mask)
	if err != nil {
		return IPv4{}, err
	}
	ip := IPv4{ipv4, bitmask}
	return ip, nil
}

func (ip IPv4) ToString() string {
	bitmask := uint32(0b11111111)
	addr := ip.addr
	octets := make([]uint32, 0)
	for i := 0; i < 4; i++ {
		octet := addr & bitmask
		octets = append(octets, octet)
		addr = addr >> 8
	}
	maskbits := bits.OnesCount32(ip.mask)
	output := fmt.Sprintf("%v.%v.%v.%v/%v", octets[3], octets[2], octets[1], octets[0], maskbits)
	return output
}

func decodeAddress(address string) (uint32, error) {
	split := strings.Split(address, ".")
	if len(split) != 4 {
		return 0, errors.New("Error decoding IPv4 address: wrong amount of octets")
	}
	var IPaddress uint32
	for i, octetstr := range split {
		segment, err := strconv.Atoi(octetstr)
		if err != nil {
			return 0, errors.Wrap(err, "Error decoding IPv4 address")
		}
		if segment > math.MaxUint8 {
			return 0, errors.New("Error decoding IPv4 address: value overflow")
		}
		// Shift octets by determined amount of bits.
		switch i {
		case 0:
			segment = segment << 24
		case 1:
			segment = segment << 16
		case 2:
			segment = segment << 8
		}
		IPaddress += uint32(segment)
	}
	return IPaddress, nil
}

func decodeMask(mask string) (uint32, error) {
	imask, err := strconv.Atoi(mask)
	var outmask uint32
	if err != nil {
		return 0, errors.Wrap(err, "Error decoding netmask")
	}
	if imask > 32 || imask < 0 {
		return 0, errors.New("Mask out of bounds")
	}
	for i := 0; i < imask; i++ {
		outmask += 1 << i
	}
	return outmask, nil
}

func main() {
	ip, err := NewIP("10.255.0.16/24")
	if err != nil {
		panic(err)
	}
	fmt.Println(ip.ToString())
}
