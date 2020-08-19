### Replace Youtube and Facebook variables
if [ -z "$YOUTUBE_URL" ] || [ -z "$YOUTUBE_KEY" ] || [ -z "$FACEBOOK_KEY" ]; then
    echo "Configuration variables missing (urls and keys)"
    exit 1
else
    VARS='$YOUTUBE_URL $YOUTUBE_KEY $FACEBOOK_KEY'
    envsubst "$VARS" < /etc/nginx/nginx.conf | tee /etc/nginx/nginx.conf > /dev/null
    
    VARS='$FACEBOOK_HOST'
    envsubst "$VARS" < /etc/nginx/nginx.conf | tee /etc/nginx/nginx.conf > /dev/null

fi