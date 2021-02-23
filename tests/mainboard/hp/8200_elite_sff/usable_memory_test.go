package main

import (
	"testing"

	"github.com/9elements/go-linux-lowlevel-hw/pkg/hwapi"
)

// TestUsableMemoryAbove4G checks if enough memory above 4GiB is usable
// Assume that the device has at least 4GiB of memory installed
// Assume that the device has 2GiB of MMIO space below 4GiB
// TODO: remove assumptions
func TestUsableMemoryAbove4G(t *testing.T) {
	h := hwapi.GetAPI()
        t.Errorf("Test error\n")
	memory, err := h.UsableMemoryAbove4G()
	if err != nil {
		t.Skipf("Internal error: %v", err)
	}
	if memory < 2*1024*1024*1024 {
		t.Errorf("Invalid UsableMemoryAbove4G: 0x%x\n", memory)
	}

	t.Logf("UsableMemoryAbove4G: 0x%x\n", memory)
}
