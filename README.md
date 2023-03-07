## OpenAI-Proxy

Golang proxy for OpenAI API

### Docker Run

```bash
docker run -d \
    --name=openai-proxy \
    --restart=always \
    -p 5333:5333 \
    -v /etc/localtime:/etc/localtime \
    yqchilde/openai-proxy:latest
```