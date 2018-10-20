#!/usr/bin/env bash

exec redis-cli KEYS "promotions:*" $REDIS_ADRESS:$REDIS_PORT| xargs redis-cli DEL $REDIS_ADRESS:$REDIS_PORT