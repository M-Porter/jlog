# jlog

Utility to format json from stdin as readible output.

## Installing

1. `go install github.com/M-Porter/jlog`
1. `dep ensure`

## Usage

```
Usage of jlog:
  -no-color
      Supresses color output

  E.g.
      tail -f error.log | jlog
```

## Example

Given the following input

```json
{"date":"2012-01-01 02:00:01", "severity":"ERROR", "msg":"Foo failed"}
{"date":"2012-01-01 02:04:02", "severity":"INFO", "msg":"Bar was successful"}
{"date":"2012-01-01 02:10:12", "severity":"DEBUG", "msg":"Baz was notified"}
```

The generated output would be

```
---------------------------------------------------------------------------------------------------------

severity: ERROR
msg: Foo failed
date: 2012-01-01 02:00:01

---------------------------------------------------------------------------------------------------------

date: 2012-01-01 02:04:02
severity: INFO
msg: Bar was successful

---------------------------------------------------------------------------------------------------------

date: 2012-01-01 02:10:12
severity: DEBUG
msg: Baz was notified
```
