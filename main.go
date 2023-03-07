package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 将所有请求路由到 handleRequest 函数进行处理
	router.Any("/*path", handleRequest)

	if err := router.Run(":5333"); err != nil {
		fmt.Println("Error starting server: ", err.Error())
	}
}

func handleRequest(c *gin.Context) {
	// 解析目标URL
	targetURL, err := url.Parse("https://api.openai.com/v1")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// 更新目标URL的主机名和协议
	targetURL.Host = "api.openai.com"
	targetURL.Scheme = "https"

	// 解析请求URL
	reqURL, err := url.Parse(c.Request.URL.String())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// 更新请求URL的主机名和协议
	reqURL.Host = targetURL.Host
	reqURL.Scheme = targetURL.Scheme

	// 复制请求头并设置 Host 和 Referer
	reqHeaders := c.Request.Header.Clone()
	reqHeaders.Set("Host", targetURL.Host)
	reqHeaders.Set("Referer", targetURL.Scheme+"://"+reqURL.Hostname())

	// 读取请求体
	reqBody, err := io.ReadAll(c.Request.Body)
	if err != nil && err != io.EOF {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// 创建 HTTP 客户端并发送代理请求
	httpClient := &http.Client{Timeout: 30 * time.Second}
	req := &http.Request{
		Method:        c.Request.Method,
		URL:           reqURL,
		Header:        reqHeaders,
		Body:          io.NopCloser(bytes.NewReader(reqBody)),
		ContentLength: int64(len(reqBody)),
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()

	// 复制响应头并设置缓存控制、跨域和安全相关的响应头
	respHeaders := resp.Header.Clone()
	respHeaders.Set("Cache-Control", "no-store")
	respHeaders.Set("Access-Control-Allow-Origin", "*")
	respHeaders.Set("Access-Control-Allow-Credentials", "true")
	respHeaders.Del("Content-Security-Policy")
	respHeaders.Del("Content-Security-Policy-Report-Only")
	respHeaders.Del("Clear-Site-Data")

	// 复制响应头到上下文的响应头
	for key, values := range respHeaders {
		for _, value := range values {
			c.Header(key, value)
		}
	}

	// 设置响应状态码
	c.Status(resp.StatusCode)

	// 复制响应体到上下文的响应体
	io.Copy(c.Writer, resp.Body)
}
