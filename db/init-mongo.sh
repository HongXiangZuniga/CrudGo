#!/bin/bash
set -e

mongoimport --db users --collection users --file /docker-entrypoint-initdb.d/dataExample.json --jsonArray
