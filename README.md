# line-bot-sdk-go-521-app

For https://github.com/line/line-bot-sdk-go/pull/521

Result: This works as expected.

[![Test](https://github.com/Yang-33/line-bot-sdk-go-521-app/actions/workflows/test.yml/badge.svg)](https://github.com/Yang-33/line-bot-sdk-go-521-app/actions/workflows/test.yml)

https://github.com/Yang-33/line-bot-sdk-go-521-lib defines http-client calling `debug.ReadBuildInfo()` internally.

This app using it, and checks if the lib shows all dependencies the app uses.

```bash
go run main.go
```
