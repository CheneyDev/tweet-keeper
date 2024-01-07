# Tweet Keeper

一个把 Tweet 转存到 Notion 的工具

### To Do

* Docker 防火墙
* 获取用户点赞的推文（所有、前一天），触发方式：快捷指令
* 手机端分享转存到 Notion，触发方式：快捷指令
* 获取用户所有媒体推文 （所有、前一天），触发方式：快捷指令

### S3 文件结构

bucket-name/user-tweets/username/YYYY/MM/DD/username-tweet-YYYY-MM-DD-XXXX.ext
bucket-name/user-likes/username/YYYY/MM/DD/username-like-YYYY-MM-DD-XXXX.ext

### Tweet 需要提取的数据

用户名，用户昵称，发布日期，推文，媒体文件

文档

请求示例

```bash
curl -X POST http://localhost:8080/getTweet \
     -H "Content-Type: application/json" \
     -d '{"tweetUrl":"https://twitter.com/example/status/1234567890", "authKey":"your_auth_key_here"}'
```