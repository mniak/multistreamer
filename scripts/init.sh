bash /scripts/vars.sh
stunnel

/auth-server/auth-server.bin &
nginx -g "daemon off;"