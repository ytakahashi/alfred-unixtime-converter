# alfred-unixtime-converter

![GO](https://img.shields.io/badge/-Go-76E1FE.svg?logo=go&style=popout)
[![Actions Status](https://github.com/ytakahashi/alfred-unixtime-converter/workflows/Go%20CI/badge.svg)](https://github.com/ytakahashi/alfred-unixtime-converter/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/ytakahashi/alfred-unixtime-converter)](https://goreportcard.com/report/github.com/ytakahashi/alfred-unixtime-converter)

Unixtime converter available as alfred workflow.

![gif image](./image/unixtime-converter.gif)

## How to use

Download a package from [release page](https://github.com/ytakahashi/alfred-unixtime-converter/releases) and install.

Activate Alfred and type `ut`.  
Input date string or unix timestamp and you see converted results.

### Acceptable input format

#### Unix Timestamp

Unix timestamp (epoch second)

- example: `1588291200`

![unixtime_s](./image/unixtime_s.png)

Unit (`s` or `ms`) can be specified (defaults to `s`)

- example: `1588291200000 ms`

![unixtime_ms](./image/unixtime_ms.png)

#### Datetime

ISO 8601 date string

- example: `2020-05-01T12:00:00Z`

![datetime](./image/datetime.png)

`yyyy-` or `yyyy-mm-ddT` can be omitted. If omitted, current year/date is applied.
