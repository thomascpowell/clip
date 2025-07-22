#!/bin/sh

if [ "$ENV" = "local" ]; then
  cp /etc/nginx/templates/default.local.conf /etc/nginx/conf.d/default.conf
else
  envsubst '$DOMAIN' < /etc/nginx/templates/default.conf.template > /etc/nginx/conf.d/default.conf
fi

exec nginx -g 'daemon off;'

