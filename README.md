# jlog

Utility to format json from stdin as readible output.

## Installing

`go install github.com/M-Porter/jlog`

## Usage

```
Usage of jlog:
  -no-color
      Supresses color output

  E.g.
      tail -f example.log | jlog
```

## Example

Given the following from stdin

```json
{ "message": "Ready to consume messages", "context": { "hostname": "docker.local.net", "application": "SearchIndexerConsumer" }, "level": 200, "level_name": "INFO", "channel": "rabbit", "datetime": { "date": "2018-12-14 17:23:56.045166", "timezone_type": 3, "timezone": "UTC" }, "extra": { "user_id": "", "logged_in_user": "", "site_id": "", "server_name": "docker.local.net", "remote_address": null, "git_branch": "live", "git_revision": "", "deployment": "", "deployment_subtype": "", "handler": "cli", "repo": "live", "http_host": "UNSET" } }
```

The generated stdout would be

```
message: Ready to consume messages
context:
    hostname: docker.local.net
    application: SearchIndexerConsumer
level: 200
level_name: INFO
channel: rabbit
datetime:
    date: 2018-12-14 17:23:56.045166
    timezone_type:
    timezone: UTC
extra:
    user_id:
    logged_in_user:
    site_id:
    server_name: docker.local.net
    remote_address: null
    git_branch: live
    git_revision:
    deployment:
    deployment_subtype:
    handler: cli
    repo: live
    http_host: UNSET
```
