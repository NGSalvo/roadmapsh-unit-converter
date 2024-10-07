#!/usr/bin/env bash

air -c ./.air.toml & \
# tailwindcss \
 # -i 'static/css/main.css' \
 # -o 'static/css/tailwind.css' \
 # --watch & \
 templ generate -watch