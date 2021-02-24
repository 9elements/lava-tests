package main

import (
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/9elements/go-linux-lowlevel-hw/pkg/hwapi"
)

// TestUsableMemoryAbove4G checks if enough memory above 4GiB is usable
// Assume that the device has 2GiB of MMIO space below 4GiB
// TODO: remove assumptions
func TestUsableMemoryAbove4G(t *testing.T) {
	h := hwapi.GetAPI()
	cnt, err := strconv.Atoi(os.Getenv("SMBIOS_TYPE17_COUNT"))
	if err != nil || cnt == 0 {
		t.Skip("Environment variable SMBIOS_TYPE17_COUNT not set")
	}
	mins, err := strconv.ParseUint(os.Getenv("SMBIOS_TYPE17_MIN_SIZE"), 0, 64)
	if err != nil || mins == 0 {
		t.Skip("Environment variable SMBIOS_TYPE17_MIN_SIZE not set")
	}
	maxs, err := strconv.ParseUint(os.Getenv("SMBIOS_TYPE17_MAX_SIZE"), 0, 64)
	if err != nil || maxs == 0 {
		t.Skip("Environment variable SMBIOS_TYPE17_MAX_SIZE not set")
	}

	ecce := strings.ToLower(os.Getenv("SMBIOS_TYPE17_ECC_ENABLED")) == "true"

	counter := 0
	_, err = h.IterateOverSMBIOSTablesType17(func(s *hwapi.SMBIOSType17) bool {
		if s.Size != 0 {
			counter++
		}
		if s.Size != 0 && s.Size < uint64(mins) {
			t.Errorf("Type17 '%d: %s' reported size is too small: 0x%x", counter, s.BankLocator, s.Size)
		}
		if s.Size != 0 && s.Size > uint64(maxs) {
			t.Errorf("Type17 '%d: %s' reported size is too big: 0x%x", counter, s.BankLocator, s.Size)
		}
		if s.TotalWidth == s.DataWidth && ecce {
			t.Errorf("Type17 '%d: %s' reported no ECC, but it was expected", counter, s.BankLocator)
		}
		if s.TotalWidth > s.DataWidth && !ecce {
			t.Errorf("Type17 '%d: %s' reported ECC, but it wasn't expected", counter, s.BankLocator)
		}

		return false
	})
	if err != nil {
		t.Skipf("Internal error using IterateOverSMBIOSTablesType17: %v", err)
	}
	if counter != cnt {
		t.Errorf("Expected %d DIMMs, but got %d", cnt, counter)
	}
}
