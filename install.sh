if [ "$USER" != "root" ]; then
  echo "You're not root!"
  echo 'Try "sudo ./install.sh" or run it as root'
  exit
fi
go build qpm.go
cp qpm /sbin/qpm
rm -rf qpm
echo done!
