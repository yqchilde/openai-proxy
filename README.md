## OpenAI-Proxy

一个十分简单的代理，用于解决 OpenAI 的网络问题。

## Usage

### Docker

```bash
docker run -d \
    --name=openai-proxy \
    --restart=always \
    -p 5333:5333 \
    -e PROXY_DOMAIN="api.openai.com" \
    yqchilde/openai-proxy:latest
```

### Vercel

[![Deploy to Vercel](https://vercel.com/button)](https://vercel.com/new/clone?repository-url=https%3A%2F%2Fgithub.com%2Fyqchilde%2Fopenai-proxy)

**设置环境变量 `PROXY_DOMAIN` 为 `api.openai.com`**

## Thanks

在一位群友帮助下，我将这个项目部署到了 Vercel 上，感谢！