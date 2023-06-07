#!/usr/bin/env bash
cd $(dirname "$0")
sqlite3 ../data.db < db.sql
# sqlite3 -csv ../data.db ".import link_data.csv URL"
