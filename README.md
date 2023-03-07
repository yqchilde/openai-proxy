## OpenAI-Proxy

一个十分简单的代理，用于解决 OpenAI 的网络问题。

### Docker Run

```bash
docker run -d \
    --name=openai-proxy \
    --restart=always \
    -p 5333:5333 \
    -v /etc/localtime:/etc/localtime \
    yqchilde/openai-proxy:latest
```