#!/usr/bin/env bash
echo "Creating tables on lake.db ..."
sqlite3 lake.db < create_tables.sql