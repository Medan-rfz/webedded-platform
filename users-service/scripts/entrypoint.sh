#!/bin/sh

CONTAINER_FIRST_STARTUP="CONTAINER_FIRST_STARTUP"
if [ ! -e /$CONTAINER_FIRST_STARTUP ]; then
    max_attempts=5
    attempt=1
    while [ $attempt -le $max_attempts ]; do
        # Migrations start
        cd /go &&
        make -s migration-up &&
        # Create admin user
        # 
        touch /$CONTAINER_FIRST_STARTUP

        if [ $? -eq 0 ]; then
            break
        else
            echo "Db reconnect... [$attempt]"
            sleep 1
        fi
        
        attempt=$((attempt+1))
    done
fi

exec app
