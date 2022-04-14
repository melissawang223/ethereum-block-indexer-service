#!/bin/bash

# Exit immediately if a command exits with a non-zero status.
set -e 

# Check location
if ( ! test -f "Makefile" ) then
    echo "Please use this script in Project Root, your are in $(pwd) now.";
    exit;
fi

# configuration file
cp env/config_general.yml env/config.yml

# install sql-migrate
go get -v github.com/rubenv/sql-migrate
go get -v github.com/rubenv/sql-migrate/...
go install github.com/rubenv/sql-migrate/...@latest

# build binary code
scripts/build.sh

# build binary code
scripts/docker.sh

# database migrate
scripts/migrate.sh
