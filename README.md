# Tweet Keeper

一个把 Tweet 转存到 Notion 的工具

### To Do

* Docker 防火墙
* 获取用户点赞的推文（所有、前一天），触发方式：快捷指令、计划任务
* 手机端分享转存到 Notion，触发方式：快捷指令
* 获取用户所有媒体推文 （所有、前一天），触发方式：快捷指令、计划任务

### S3 文件结构

bucket-name/user-tweets/username/YYYY/MM/DD/username-tweet-YYYY-MM-DD-XXXX.ext
bucket-name/user-likes/username/YYYY/MM/DD/username-like-YYYY-MM-DD-XXXX.ext

### Tweet 需要提取的数据

用户名，用户昵称，发布日期，推文，媒体文件

### 请求示例

/getTweet

```bash
curl -X POST http://localhost:8080/getTweet \
     -H "Content-Type: application/json" \
     -d '{"tweetUrl":"https://twitter.com/example/status/1234567890", "authKey":"your_auth_key_here"}'
```

### 原始 API 响应示例

仅文字 tweet

```
{
  "likes": 107774,
  "created_at": "Sat Jan 06 21:40:25 +0000 2024",
  "status": "active",
  "text": "The US construction industry can increase available housing by 1% to 2% per year. \n\nBut the US population is only 330M out of 8B. \n\nSo if just ~4% of Earth moves here, housing would need to double, which is impossible, causing a massive homeless problem and house prices to be astronomically unaffordable.",
  "display_text": "The US construction industry can increase available housing by 1% to 2% per year. \n\nBut the US population is only 330M out of 8B. \n\nSo if just ~4% of Earth moves here, housing would need to double, which is impossible, causing a massive homeless problem and house prices to be ",
  "retweets": 20109,
  "bookmarks": 4270,
  "quotes": 1686,
  "reply_to": null,
  "replies": 12868,
  "lang": "en",
  "views": "23388974",
  "conversation_id": "1743749384389443975",
  "author": {
    "rest_id": "44196397",
    "name": "Elon Musk",
    "screen_name": "elonmusk",
    "image": "https://pbs.twimg.com/profile_images/1683325380441128960/yRsRRjGO_normal.jpg",
    "blue_verified": true,
    "sub_count": 168435894
  },
  "media": null,
  "id": "1743749384389443975"
}
```

