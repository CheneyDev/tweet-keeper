# Tweet Keeper

一个把 Tweet 转发到 Telegram 的工具

### To Do

Docker 防火墙问题

文档

请求示例

```bash
curl -X POST http://localhost:8080/getTweet \
     -H "Content-Type: application/json" \
     -d '{"tweetUrl":"https://twitter.com/example/status/1234567890", "authKey":"your_auth_key_here"}'
```