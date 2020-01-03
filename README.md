# AWS Lambda Golang Telegram Bot

## This bot will reply you with your message

An environment variable required:
```bash
export T_TOKEN=<your_telegram_bot_token>
```

Set a webhook:
```bash
curl --request POST \
  --url https://api.telegram.org/bot<your_telegram_bot_token>/setWebhook \
  --header 'content-type: application/json' \
  --data '{
	"url": "<aws_api_gateway_invoke_url>"
}'
```