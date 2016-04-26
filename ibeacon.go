package ibeacon

import (
	"bytes"
	"fmt"

	"github.com/paypal/gatt"
)

func isIbeacon(a *gatt.Advertisement) bool {
	if bytes.Compare(a.ManufacturerData[0:4], []byte{76, 0, 2, 21}) == 0 {
		return true
	}
	return false
}

func getUuid(a *gatt.Advertisement) string {
	uuid10 := a.ManufacturerData[4:20]
	uuid16 := fmt.Sprintf("%x", uuid10)
	return uuid16
}

func getMajor(a *gatt.Advertisement) string {
	major10 := a.ManufacturerData[20:22]
	major16 := fmt.Sprintf("%x", major10)
	return major16
}

func getMinor(a *gatt.Advertisement) string {
	minor10 := a.ManufacturerData[22:24]
	minor16 := fmt.Sprintf("%x", minor10)
	return minor16
}
