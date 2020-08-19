### Replace Youtube variables
if [ -z "$YOUTUBE_URL" ] || [ -z "$YOUTUBE_KEY" ] || [ -z "$FACEBOOK_URL" ] || [ -z "$FACEBOOK_KEY" ]
then
    echo "Configuration variables missing"
    exit 1
else
    VARS='$YOUTUBE_URL $YOUTUBE_KEY $FACEBOOK_URL $FACEBOOK_KEY'
    envsubst "$VARS" </etc/nginx/nginx.conf | tee /etc/nginx/nginx.conf >/dev/null
fi