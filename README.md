# Usage
1. Create Slack App at [https://api.slack.com/apps](https://api.slack.com/apps).
2. Add `emoji:read` to App Permission Scope.
3. Set environment variable `$SLACK_API_TOKEN`.
4. Simply run `build/slack-emoji-export_{YOUR_ARCHITECTURE}` and emojis will be generated at `emoji` directory.

# Development
1. Install `golang/dep` and `mitchellh/gox`.
2. `$ dep ensure`.
3. `$ go run main.go`.
4. To rebuild binaries, run `$ make x-build`.

