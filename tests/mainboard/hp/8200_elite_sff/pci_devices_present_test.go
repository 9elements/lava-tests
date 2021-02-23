package main

import (
	"testing"

	"github.com/9elements/go-linux-lowlevel-hw/pkg/hwapi"
)

// TestPCIDevicesPresent checks if all required PCI devices are visible
func TestPCIDevicesPresent(t *testing.T) {
	h := hwapi.GetAPI()

	l := []hwapi.PCIDevice{}

	err := h.PCIEnumerateVisibleDevices(func(d hwapi.PCIDevice) (abort bool) {
		l = append(l, d)
		return false
	})
	if err != nil {
		t.Skipf("Internal error: %v", err)
	}

	reference := []hwapi.PCIDevice{
		{Bus: 0, Device: 0, Function: 0},
		{Bus: 0, Device: 0x19, Function: 0},
		{Bus: 0, Device: 0x1a, Function: 0},
		{Bus: 0, Device: 0x1b, Function: 0},
		{Bus: 0, Device: 0x1c, Function: 0},
		{Bus: 0, Device: 0x1c, Function: 1},
		{Bus: 0, Device: 0x1c, Function: 4},
		{Bus: 0, Device: 0x1c, Function: 6},
		{Bus: 0, Device: 0x1d, Function: 0},
		{Bus: 0, Device: 0x1e, Function: 0},
		{Bus: 0, Device: 0x1f, Function: 0},
		{Bus: 0, Device: 0x1f, Function: 2},
		{Bus: 0, Device: 0x1f, Function: 3},
	}

	for _, r := range reference {
		found := false
		for i := range l {
			if l[i].Device == r.Device &&
				l[i].Function == r.Function &&
				l[i].Bus == r.Bus {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("PCI device %#v is missing\n", r)
		}
	}
	t.Logf("Found all PCI devices\n")
}
