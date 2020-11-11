# Data Source `timestamp`

The timestamp data source allows you to converts a local timestamp (e.g. without timezone info) into one with the specified timezone and layout.

## Example Usages

### Use local timezone

```terraform
# 2020/11/22 11:11:11 -0800
data "localtime_timestamp" "local_timezone" {
  local_time = "2020/11/22 11:11:11"
}
```

### Specify a timezone location

```terraform
# 2020/11/22 11:11:11 -0500
data "localtime_timestamp" "specific_timezone" {
  local_time = "2020/11/22 11:11:11"
  location   = "America/New_York"
}
```

### Customize only the timezone part of layout

```terraform
# 2020/11/22 11:11:11 PST
data "localtime_timestamp" "custom_timezone_layout" {
  local_time      = "2020/11/22 11:11:11"
  timezone_layout = " MST"
}
```

### Customize timestamp layout

```terraform
# 2020-11-22T11:11:11Z08:00
data "localtime_timestamp" "custom_layout" {
  local_time = "2020/11/22 11:11:11"
  layout     = "2006-01-02T15:04:05Z07:00"
}
```

## Argument Reference

- `local_time` - (Required) The local timestamp without timezone info. It can be in any valid layout.
- `location` - (Optional) The location of the timezone where the exported timestamp will be. It must be a valid [tz database name](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List). Default to `Local`.
- `layout` - (Optional) The layout that the exported timestamp will use. It must follow the [Golang format](https://golang.org/pkg/time/#pkg-constants). If this parameter is provided, `timezone_layout` will be ignored. Otherwise the exported timestamp will use the same layout as in `local_time`, plus `timezone_layout`.
- `timezone_layout` - (Optional) The timezone part of `layout`. It's useless when `layout` is provided. Default to " -0700".

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `timestamp` - The timestamp with the specific timezone and layout.
