#!/usr/bin/env bash

air -c ./.air.toml & \
npx tailwindcss \
 -i 'static/styles/main.css' \
 -o 'static/styles/style.css' \
 --watch & \
 templ generate -watch

 # templ generate -watch -proxy="http://localhost:3000" -cmd="go run main.go"