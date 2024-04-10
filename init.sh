#!/bin/bash
mongod --bind_ip_all &
sleep 5
mongoimport --host mongodb --db users --collection users --type json --file /dataExample.json --jsonArray
tail -f /dev/null
