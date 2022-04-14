#!/bin/bash

# Exit immediately if a command exits with a non-zero status.
set -e 

# Check location
if ( ! test -f "Makefile" ) then
    echo "Please use this script in Project Root, your are in $(pwd) now.";
    exit;
fi

echo "[`date "+%Y/%m/%d %H:%M:%S"`] Start Build.";

go build .

echo "[`date "+%Y/%m/%d %H:%M:%S"`] Finished.";