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

### Customize only the timezone part of layout

```terraform
# 2020/11/22 11:11:11 PST
data "localtime_timestamp" "custom_timezone_layout" {
  local_time      = "2020/11/22 11:11:11"
  layout_timezone = "MST"
}
```

## Argument Reference

- `local_time` - (Required) The local timestamp without timezone info. Most time layouts are supported. See the full list [here](https://github.com/araddon/dateparse/blob/8aadafed4dc4aee1363ec2a04c9c954544ee54dc/example/main.go#L12-L111).
- `location` - (Optional) The location of the timezone where the exported timestamp will be. It must be a valid [tz database name](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List). If this parameter is not provided, it will use your local timezone.
- `layout` - (Optional) The layout that the exported timestamp will use. It must follow the [Golang format](https://golang.org/pkg/time/#pkg-constants). This parameter conflicts with `layout_timezone`. If this parameter is not provided, the exported timestamp will use the same layout as in `local_time`.
- `layout_timezone` - (Optional) The timezone part of a [layout](https://golang.org/pkg/time/#pkg-constants), e.g. `-07:00` and `MST`. Default to "-0700". It conflicts with `layout`, which specifies a full layout including the timezone part.

## Attributes Reference

In addition to all the arguments above, the following attributes are exported.

- `timestamp` - The timestamp with the specific timezone and layout.
