FROM nginx:alpine

COPY default.conf.template /etc/nginx/templates/default.conf.template
COPY default.local.conf /etc/nginx/templates/default.local.conf
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]

