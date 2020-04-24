# coreboot payloads for firmware testing

The payloads will be stitched into the ROM before it's
being send to the DUT.

Update with care.

Every binary has a matching YAML config file.

## Precompiled binaries

*SeaBIOS*: Used on X86
*Tianocore*: Used on X86
*Linux+BusyBox*: Used on X86

## YAML contents

**Example:**

```
name: "SeaBIOS"
commithash: "f21b5a4aeb020f2a5e2c6503f906a9349dd2f069"
repo_url: "https://review.coreboot.org/seabios.git"
version: "1.13.0"
filename: "seabios.elf"
license: "GPLv2"
cbfstool_args:
        - "-f"
        - "%PAYLOAD%"
        - "-n"
        - "fallback/payload"
        - "-c"
        - "lzma"
whitelist:
        - "CONFIG_ARCH_X86=y"
        - "CONFIG_TTYS0_BASE=0x3f8"
blacklist:
        - "CONFIG_GDB_WAIT=y"
harddisk_image: https://blobs.9esec.io/nightly/Fedora-HWT-disk-30-minimal-MBR.img.xz
harddisk_image_compression: "xz"

kernel_image: https://blobs.9esec.io/os/nightly/vmlinuz-5.5.17-200.fc31.x86_64
kernel_type: zimage
ramdisk_image: https://blobs.9esec.io/os/nightly/initramfs-5.5.17-200.fc31.x86_64.img
ramdisk_compression: xz
dtb_image: https://blobs.9esec.io/os/nightly/live-rootfs.squashfs.img
```

**Description:**

* `filename` configures the path to the payload binary.
* `filename2` configures the path to another the payload binary.
* `kernel_image` a kernel used for netboot only
* `kernel_type` passed to some bootloaders
* `ramdisk_image` a initramfs used for netboot only
* `ramdisk_compression` the compression of the ramdisk (usually xz)
* `dtb_image` a devicetree or squashfs
* `license` the license of the project and the binary.
* `whitelist` defines Kconfig symbols that must be present (like serial)
* `blacklist` defines Kconfig symbols that must not be present (like GDB wait)
* `harddisk_image` defines the harddisk to use (optional)
* `repo_url` The URL to the repository used (optional)
* `commithash` The commit hash when the payload was build (optional)
* `version` defines the payload internal versioning
* In `cbfstool` args:

** `%PAYLOAD%` will be substituted with `filename`
** `%INITRD%` will be substituted with `filename2`
