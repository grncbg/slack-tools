# slack-tools

## Getting Started

### Install

`go get -u github.com/grncbg/slack-tools`

To make it more useful, you can add `$GOPATH/bin` to PATH environment variable.

### Configure

Config file is `$HOME/.slack-tools.(yaml|toml|json|env)`.
And set slack token to config file.
  
#### Example
~/.slack-tools.toml
```TOML
token = SLACK_TOKEN
```

### Useage

#### Invite

- Invite a user to channel:
  `slack-tools invite -c <channel_id> -u <user_id>`

- Invite users to channel:
  `slack-tools invite -c <channel_id> -l <user_list_file>`

#### Get user ID

`slack-tools ids [-b]`
- `-b`: with bot user ID
