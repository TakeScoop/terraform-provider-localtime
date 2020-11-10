terraform {
  required_providers {
    localtime = {
      source  = "takescoop.jfrog.io/localtime"
      version = "~> 0.1"
    }
  }
}

provider "localtime" {}

# data.localtime_timestamp.same_layout.timestamp == "2020/11/22 11:11:11 -0800"
data "localtime_timestamp" "same_layout" {
  local_time = "2020/11/22 11:11:11"
  location   = "America/Los_Angeles"
}

# data.localtime_timestamp.same_layout.timestamp == "2020/11/22 11:11:11 PST"
data "localtime_timestamp" "custom_timezone_layout" {
  local_time      = "2020/11/22 11:11:11"
  location        = "America/Los_Angeles"
  timezone_layout = " MST"
}

# data.localtime_timestamp.same_layout.timestamp == "2020-11-22T11:11:11Z08:00"
data "localtime_timestamp" "custom_layout" {
  local_time = "2020/11/22 11:11:11"
  location   = "America/Los_Angeles"
  layout     = "2006-01-02T15:04:05Z07:00"
}

# data.localtime_timestamp.same_layout.timestamp == "Mon, 22 Nov 2020 11:11:11 -0800"
data "localtime_timestamp" "custom_predefined_layout" {
  local_time = "2020/11/22 11:11:11"
  location   = "America/Los_Angeles"
  layout     = "RFC1123Z"
}
