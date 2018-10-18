#!/usr/bin/env bash

exec redis-cli KEYS "promotions:*" | xargs redis-cli DEL