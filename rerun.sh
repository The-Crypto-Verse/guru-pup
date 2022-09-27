go build .
sudo systemctl restart gurupup.service
journalctl -f -u gurupup.service