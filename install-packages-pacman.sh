#!/bin/bash

if [ "$EUID" -ne 0 ]
    then echo "must run as root"
        exit
fi

echo "=======> Downloading official packages"
echo

pacman -Sy --needed $(<packages-pacman.txt)
