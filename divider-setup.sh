#!/usr/bin/env bash
read -p 'DSP Microservice Address: ' address

DSP_MICROSERVICE_ADDRESS="DSP_MICROSERVICE_ADDRESS=$address"
grep -qF "$DSP_MICROSERVICE_ADDRESS" /etc/environment || echo "$DSP_MICROSERVICE_ADDRESS" >> /etc/environment

mkdir /etc/divider-sensors

curl -o /etc/divider-sensors/config.json https://raw.githubusercontent.com/byuoitav/divider-sensors-microservice/master/config.json
vi /etc/divider-sensors/config.json

curl -o /etc/systemd/system/divider-sensors.service https://raw.githubusercontent.com/byuoitav/divider-sensors-microservice/master/divider-sensors.service
wget  https://github.com/byuoitav/divider-sensors-microservice/releases/download/v0.3.1/divider-sensors-microservice
chmod 755 divider-sensors-microservice
mv divider-sensors-microservice /etc/divider-sensors

systemctl daemon-reload
systemctl disable salt-minion.service
systemctl enable divider-sensors.service
systemctl start divider-sensors.service
