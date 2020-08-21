![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/mniak/multistreamer?style=flat-square)

mniak/multistreamer
=====================

RMTP server that restreams to Youtube and Facebook

## Usage
```bash
docker run -d --name myserver \
  --env YOUTUBE_KEY=<Put the Youtube Key here> \
  --env FACEBOOK_KEY=<Put the Facebook Key here> \
  --env STREAMING_KEY=<Put the streaming key here> \
  --p 1935:1935 \
  mniak/multistreamer
```

### Required environment variables
#### `YOUTUBE_KEY`
The youtube streaming key

#### `FACEBOOK_KEY`
The facebook streaming

#### `STREAMING_KEY`
Your own streaming key, the one that will authenticate you in this container.

## Features/Roadmap
- [x] Stream to YouTube
- [x] Stream to Facebook
- [x] Authentication
- [ ] Retry logic for YouTube
- [ ] Retry logic for Facebook
- [ ] Option to disable YouTube streaming
- [ ] Option to disable Facebook streaming
- [ ] Receive streaming keys on request, avoiding need to restart when the streaming key changes
