package localtime

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

type TestAccDataSourceTimestampCase struct {
	Name              string
	Config            string
	ExpectedTimestamp string
}

func TestAccDataSourceTimestamp_basic(t *testing.T) {
	testAccDataSourceTimestampHelper(t, &TestAccDataSourceTimestampCase{
		Name:              "basic",
		ExpectedTimestamp: "2020/11/22 11:11:11 -0800",
		Config: `
		data "localtime_timestamp" "basic" {
		  local_time = "2020/11/22 11:11:11"
		  location   = "America/Los_Angeles"
		}`,
	})
}

func TestAccDataSourceTimestamp_layout(t *testing.T) {
	testAccDataSourceTimestampHelper(t, &TestAccDataSourceTimestampCase{
		Name:              "layout",
		ExpectedTimestamp: "2020-11-22T11:11:11 PST",
		Config: `
		data "localtime_timestamp" "layout" {
		  local_time = "2020/11/22 11:11:11"
		  location   = "America/Los_Angeles"
		  layout     = "2006-01-02T15:04:05 MST"
		}`,
	})
}

func TestAccDataSourceTimestamp_timezone(t *testing.T) {
	testAccDataSourceTimestampHelper(t, &TestAccDataSourceTimestampCase{
		Name:              "timezone",
		ExpectedTimestamp: "2020/11/22 11:11:11 PST",
		Config: `
		data "localtime_timestamp" "timezone" {
		  local_time      = "2020/11/22 11:11:11"
		  location        = "America/Los_Angeles"
		  layout_timezone = "MST"
		}`,
	})
}

func testAccDataSourceTimestampHelper(t *testing.T, test_case *TestAccDataSourceTimestampCase) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: test_case.Config,
				Check: resource.TestCheckResourceAttr(
					fmt.Sprintf("data.localtime_timestamp.%s", test_case.Name),
					"timestamp",
					test_case.ExpectedTimestamp,
				),
			},
		},
	})
}
