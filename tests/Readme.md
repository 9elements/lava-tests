# master configuration file

**Define all tests that should be included into the LAVA job here.**

## Format
The file format is YAML.

The syntax follows LAVA test definition syntax, but was extended by `whitelist`
and `blacklist`.

All whitelist entries must be present in the boards Kconfig, otherwise the
test isn't run.

No blacklist entry must be present in the boards Kconfig, otherwise the
test isn't run.


Example:
```
- repository: https://github.com/9elements/lava-tests.git
  from: git
  path: tests/x86/acpi-fwts-tests.yaml
  name: acpi-fwts-test
  whitelist:
    CONFIG_HAVE_ACPI_TABLES : y
  blacklist:
# QEMU provides it's own ACPI, don't test it
    CONFIG_BOARD_EMULATION_QEMU_X86 : y
    CONFIG_BOARD_EMULATION_QEMU_X86_I440FX : y
    CONFIG_BOARD_EMULATION_QEMU_X86_Q35 : y
```

