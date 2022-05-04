# docker-redirect

Simple image to redirect to another website

## usage

```bash
docker run --detach --restart=always --net=host --env=PORT=80 --env=REDIRECT_TO=http://example.com ghcr.io/chneau/docker-redirect
```
