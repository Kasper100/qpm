echo 'export PATH="$HOME/.qpm:$PATH"' >> ~/.bashrc
echo "Exported .qpm to PATH"
sudo pacman -S go
./build.sh
