# slack-tools

## 入門

### インストール

`go get -u github.com/grncbg/slack-tools`

もし`$GOPATH/bin`がPATHに含まれていないのであれば、追加することで使いやすくなります。  
以下、PATHが設定されている前提で記載します。

### 設定

Viperでサポートされるフォーマットが利用可能です。  
`$HOME/.slack-tools.(yaml|toml|json|env)`等にSlackのtokenを設定してください。
  
#### 例
~/.slack-tools.toml
```TOML
token = "SLACK_TOKEN"
```

### 使い方

#### ユーザをチャンネルに招待する

- ユーザ一人をチャンネルに追加する:
  `slack-tools invite -c <channel_id> -u <user_id>`

- 複数のユーザをチャンネルに追加する:
  `slack-tools invite -c <channel_id> -l <user_list_file>`

#### ユーザIDの取得

`slack-tools ids [-b]`
- `-b`: ボットのIDも表示する
