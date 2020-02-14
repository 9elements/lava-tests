## Prebuild coreboot ROMs for periodic health checks

The ROMs only contain a valid BIOS partition. The flasher has to
provide other binaries or has to flash only the BIOS partition.

By now the ROMs are build by hand and manually tested.

## Naming conventions

1. coreboot ROMs start with `coreboot-`
2. followed by the board name, whithout CONFIG_ prefix
3. followed by the payload name
4. followed by the payload version

**Example:**

	coreboot-BOARD_HP_COMPAQ_HP8200-Seabios-1.13.0.rom
