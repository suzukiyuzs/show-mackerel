# show-mackerel

## Description

Show mackerel monitor setting and agent checkplugin settings by markdown format.  

caution: Host type monitor setting only can be parsed. Service type and external type can't be parsed. 

## Build

```
$ go get github.com/suzukiyuzs/show-mackerel
$ cd $GOPATH/src/github.com/suzukiyuzs/show-mackerel
$ go build .
```

## Help


```
$ ./show-mackerel -h
Usage:
  show-mackerel [OPTIONS]

Application Options:
  -m, --mackerel= Mackerel type(check or monitor)
  -f, --file=     Mackerel file path

Help Options:
  -h, --help      Show this help message

$
```

- Mackerel type option:
  - check: Show mackerel agent checkplugin setting
  - monitor: Show mackerel monitor setting

## Example

### Show agent checkplugin settings

```
$ cat ./check-pluguin
[plugin.checks.ssh]
command = "ruby /path/to/check-ssh.rb"
notification_interval = 60
max_check_attempts = 1
check_interval = 5
timeout_seconds = 45
prevent_alert_auto_close = true
env = { HOST = "hostname", PORT = "port" }
action = { command = "ruby /path/to/notify_something.rb", env = { NOTIFY_API_KEY = "API_KEY" } }
memo = "This check monitor is ..."
$
$ ./show-mackerel -m check -f ./check-pluguin
# Mackerel check plugin
| Name |          Command           | NotificationInterval | MaxCheckAttempts | CheckInterval | TimeoutSeconds | PreventAlertAutoClose |           Env            |          Action.Command           | Action.TimeoutSeconds |       Action.Env       |           Memo            |
|------|----------------------------|----------------------|------------------|---------------|----------------|-----------------------|--------------------------|-----------------------------------|-----------------------|------------------------|---------------------------|
| ssh  | ruby /path/to/check-ssh.rb |                   60 |                1 |             5 | 45s            | true                  | HOST=hostname, PORT=port | ruby /path/to/notify_something.rb | 0s                    | NOTIFY_API_KEY=API_KEY | This check monitor is ... |
$
```


### Show monitor settings

```
$ export MACKEREL_APIKEY=<Mackerel API Key>
$ mkr monitors pull
$ cat monitors
{
    "monitors": [
        {
            "id": "3pd4qrKKhRw",
            "id": "0123456789a",
            "name": "disk.aa-00.writes.delta",
            "memo": "This monitor is ...",
            "type": "host",
            "duration": 3,
            "metric": "disk.aa-00.writes.delta",
            "operator": ">",
            "warning": 20000.0,
            "critical": 400000.0,
            "maxCheckAttempts": 3,
            "notificationInterval": 60,
            "scopes": [
                "Hatena-Blog"
            ],
            "excludeScopes": [
                "Hatena-Bookmark: db-master"
            ]
        }
    ]
}
$
$ ./show-mackerel -m monitor -f ./monitors.json
# Mackerel monitor
|     ID      |          Name           |        Memo         | Type | IsMute | NotificationInterval |         Metric          | Operator | Warning | Critical | Duratoin | MaxCheckAttempts |   Scopes    |       ExcludeScopes        |
|-------------|-------------------------|---------------------|------|--------|----------------------|-------------------------|----------|---------|----------|----------|------------------|-------------|----------------------------|
| 0123456789a | disk.aa-00.writes.delta | This monitor is ... | host | false  |                   60 | disk.aa-00.writes.delta | >        |   20000 |   400000 |        3 |                3 | Hatena-Blog | Hatena-Bookmark: db-master |
$
```

