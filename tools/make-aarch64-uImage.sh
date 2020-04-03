#!/bin/bash

[ "$#" -ne 3 ] && echo "run: $0 <kernel> <initrd> <dtb>" && exit 1

mkdir -p ./tmp
xz -c -k -f --format=lzma --lzma1=dict=1MiB,lc=3,lp=0,pb=3 < $1 > ./tmp/kernel.lzma
xz -k -f --check=crc32 --lzma2=dict=512KiB < $2 > ./tmp/initramfs.cpio.xz
cp $3 ./tmp/file.dtb

mkimage -f aarch64.its uImage

rm -rf ./tmp
