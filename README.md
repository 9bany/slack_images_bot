# Slack Images Bot
Slack bot send a relaxing image like a beauty girl
## Commands
|  Name | Command  | Description  | Already  |
|---|---|---|---|
| Relaxing  | `relax`  | Will send a relaxing message | YES  |
## Preparing your Slack App

To use bot_images you'll need to create a Slack App, either manually or with an app manifest.

Please read this: https://github.com/shomali11/slacker#preparing-your-slack-app

## Crawl data
- Down image datas to `data/images` folder
- Use `scripts/main.go` to upload them to db
- Tools 
    - [instaloader](https://github.com/instaloader/instaloader)
```
$ instaloader profile <instagram_username>  --no-captions --no-videos --no-video-thumbnails --no-metadata-json
```
## Docker

- Env file
```
DB_DRIVER=postgres 
DB_NAME='images_demo' 
DB_SOURCE='postgresql://user:pass@localhost:port/
dbname' 
SLACK_APP_TOKEN=''
SLACK_BOT_TOKEN='' 
```
Command
```
$ docker run --env-file ./.env ghcr.io/9bany/bot_workflows
```
## Contributing

If you want to colaborate check the project's issues.

1. Fork the repository
2. Implement your solution
3. Commit
4. Open a Pull Request

Thanks!
