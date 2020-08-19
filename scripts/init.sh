bash /scripts/vars.sh && \
/auth-server/auth-server.bin && \
stunnel && \
nginx -g "daemon off;"