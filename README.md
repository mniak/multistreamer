## mniak/multistreamer

RMTP server that restreams to Youtube and Facebook


Usage
```bash
docker run -d --name myserver \
  --env YOUTUBE_URL=<Put the Youtube URL here> \
  --env YOUTUBE_KEY=<Put the Youtube Key here> \
  --env FACEBOOK_URL=<Put the Facebook URL here> \
  --env FACEBOOK_KEY=<Put the Facebook Key here> \
  --port 1935:1935 \
  mniak/multistreaming
```