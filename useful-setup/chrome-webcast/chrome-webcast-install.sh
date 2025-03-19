#!/bin/bash

sudo apt install -y xvfb xserver-xorg-video-dummy xbase-clients libutempter0
sudo apt install -y python3-packaging python3-psutil
# Check if wget is installed
if ! command -v wget &> /dev/null; then
    echo "wget is not installed. Installing now..."
    sudo apt update
    sudo apt install wget -y
    echo "wget has been installed successfully."
else
    echo "wget is already installed."
fi


wget -P ~/ https://dl.google.com/linux/direct/chrome-remote-desktop_current_amd64.deb

sudo dpkg -i ~/chrome-remote-desktop_current_amd64.deb

rm ~/chrome-remote-desktop_current_amd64.deb

echo "======== Successfully installed chrome remote desktop ==========="


