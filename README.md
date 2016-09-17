# Slack-go-webhook
A Simple docker tool to push message to slack via webhook

# Usage
```shell
docker pull ngtrieuvi92/slack-webhook-go
docker run -e SLACK_HOOK_URL="your_slack_webhook_rul" \
  -e SLACK_USERNAME="CI-bot" \
  -e SLACK_CHANNEL="#general" \
  -e SLACK_TEXT_MESSAGE="Build completed" \
  ngtrieuvi92/slack-webhook-go
```
