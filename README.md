## mniak/multistreamer

RMTP server that restreams to Youtube and Facebook


Usage
```bash
docker run -d --name myserver \
  --env YOUTUBE_URL=<Put the Youtube URL here> \
  --env YOUTUBE_KEY=<Put the Youtube Key here> \
  --env FACEBOOK_KEY=<Put the Facebook Key here> \
  --env STREAMING_KEY=<Put the streaming key here> \
  --port 1935:1935 \
  mniak/multistreamer
```

### `YOUTUBE_URL`
The youtube streaming URL

### `YOUTUBE_KEY`
The youtube streaming key

### `FACEBOOK_KEY`
The facebook streaming

### `STREAMING_KEY`
Your own streaming key, the one that will authenticate you in this container.