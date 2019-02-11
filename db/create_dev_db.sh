#!/usr/bin/env bash
echo "Creating tables on lake.db ..."
sqlite3 lake.db < create_tables.sql

echo "Populating tables on lake.db ..."
sqlite3 lake.db < populate_tables.sql
