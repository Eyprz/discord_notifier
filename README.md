# discord_notifier
This is software to send notifications to Discord via webhook

## How to use?

`./discord_notifier [-w <webhook_url>] [-a <avatar_url>] [-u <username>] <message>`

The `message` argument is required. By default, the contents of `config.yaml` are used, and if options are specified, the request will be sent with the config overridden only where applicable.

## Example

```bash
./discord_notifier -u "Your Bot" "Hi! This is messaging you from the command line!"
```

## Configuration

The configuration file is `config.yaml` and is located in the `config` directory. The default configuration is as follows:

```yaml
webhook_url: <Your discord webhook url>
avatar_url: <Your avatar url>
username: <Your bot name>
```
