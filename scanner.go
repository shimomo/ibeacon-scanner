package ibeacon

import (
	"log"
	"time"

	"github.com/paypal/gatt"
	"github.com/paypal/gatt/examples/option"
)

func Scan(discover func(string, string, string, int), sleep int) {
	onPeriphDiscovered := func(p gatt.Peripheral, a *gatt.Advertisement, rssi int) {
		if len(a.ManufacturerData) == 0 {
			return
		}

		if isIbeacon(a) {
			discover(getUuid(a), getMajor(a), getMinor(a), rssi)
		}
	}

	onStateChanged := func(d gatt.Device, s gatt.State) {
		switch s {
		case gatt.StatePoweredOn:
			for {
				d.Scan([]gatt.UUID{}, false)
				time.Sleep(time.Duration(sleep) * time.Millisecond)
				d.StopScanning()
			}
		default:
			d.StopScanning()
		}
	}

	d, err := gatt.NewDevice(option.DefaultClientOptions...)
	if err != nil {
		log.Fatalf("Failed to open device, err: %s\n", err)
		return
	}

	d.Handle(gatt.PeripheralDiscovered(onPeriphDiscovered))
	d.Init(onStateChanged)
	select {}
}
