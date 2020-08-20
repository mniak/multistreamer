env
/tools/tools.bin mkconfig nginx -o /etc/nginx/nginx.conf
stunnel

/tools/tools.bin server &
nginx -g "daemon off;"