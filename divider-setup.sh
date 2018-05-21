#!/usr/bin/env bash
read -p 'DSP Microservice Address: ' address1

read -p 'Event Router Address: ' address2

DSP_MICROSERVICE_ADDRESS="DSP_MICROSERVICE_ADDRESS=$address1"
LOCAL_ENVIRONMENT="LOCAL_ENVIRONMENT=true"
EVENT_ROUTER_ADDRESS="EVENT_ROUTER_ADDRESS=$address2"
CONTACTS_CONFIG_FILE="CONTACTS_CONFIG_FILE=/etc/divider-sensors/config.json"

grep -qF "$DSP_MICROSERVICE_ADDRESS" /etc/environment || echo "$DSP_MICROSERVICE_ADDRESS" >> /etc/environment
grep -qF "$LOCAL_ENVIRONMENT" /etc/environment || echo "$LOCAL_ENVIRONMENT" >> /etc/environment
grep -qF "$EVENT_ROUTER_ADDRESS" /etc/environment || echo "$EVENT_ROUTER_ADDRESS" >> /etc/environment
grep -qF "$CONTACTS_CONFIG_FILE" /etc/environment || echo "$CONTACTS_CONFIG_FILE" >> /etc/environment

mkdir /etc/divider-sensors

curl -o /etc/divider-sensors/config.json https://raw.githubusercontent.com/byuoitav/divider-sensors-microservice/master/config.json
vi /etc/divider-sensors/config.json

curl -o /etc/systemd/system/divider-sensors.service https://raw.githubusercontent.com/byuoitav/divider-sensors-microservice/master/divider-sensors.service
wget  https://github.com/byuoitav/divider-sensors-microservice/releases/download/v0.4/divider-sensors-microservice
chmod 755 divider-sensors-microservice
mv divider-sensors-microservice /etc/divider-sensors

systemctl daemon-reload
systemctl disable salt-minion.service
systemctl enable divider-sensors.service
systemctl start divider-sensors.service
