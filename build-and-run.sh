#!/usr/bin/bash

docker compose up
mkdir -p ~/sanepar-falta-agua
cp ./.env.example ~/sanepar-falta-agua/.env
cp ./dist/sanepar-falta-agua ~/sanepar-falta-agua/sanepar-falta-agua
cd ~/sanepar-falta-agua
chmod +x ./sanepar-falta-agua

./sanepar-falta-agua