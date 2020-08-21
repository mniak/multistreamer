# Generate configuration files
/tools/tools.bin mkconfig nginx -o /etc/nginx/nginx.conf
/tools/tools.bin mkconfig stunnel -o /etc/stunnel/stunnel.conf

# Start STunnel4
stunnel

# Start tools server
/tools/tools.bin server &

# Start NGINX
nginx -g "daemon off;"