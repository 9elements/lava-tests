package main

import (
	"testing"

	"github.com/9elements/go-linux-lowlevel-hw/pkg/hwapi"
)

// TestUsableMemoryAbove4G checks if enough memory above 4GiB is usable
// Assume that the device has 2GiB of MMIO space below 4GiB
// TODO: remove assumptions
func TestUsableMemoryAbove4G(t *testing.T) {
	h := hwapi.GetAPI()
	memoryInstalled := uint64(0)
	_, err := h.IterateOverSMBIOSTablesType17(func(s *hwapi.SMBIOSType17) bool {
		memoryInstalled += s.Size
		return false
	})
	if err != nil || memoryInstalled == 0 {
		// Assume that the device has at least 4GiB of memory installed
		t.Logf("SMBIOS type 17 access error: %v\n", err)
		memoryInstalled = 0x100000000
	} else {
		t.Logf("SMBIOS reports 0x%x bytes of memory\n", memoryInstalled)
	}
	if memoryInstalled < 0x100000000 {
		// FIXME: There still might be memory mapped above 4GiB
		t.Skip("Not enough memory installed")
	}
	memory, err := h.UsableMemoryAbove4G()
	if err != nil {
		t.Skipf("Internal error: %v", err)
	}
	if memory < 2*1024*1024*1024 {
		t.Errorf("Invalid UsableMemoryAbove4G: 0x%x\n", memory)
	}

	t.Logf("UsableMemoryAbove4G: 0x%x\n", memory)
}
