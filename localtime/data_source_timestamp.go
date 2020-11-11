package hashicups

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTimestamp() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTimestampRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"layout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"local_time": &schema.Schema{
				Type: schema.TypeString,
			},
			"location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"timestamp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"timezone_layout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  " -0700",
			},
		},
	}
}

func dataSourceTimestampRead(ctx context.Context, d *schema.ResourceData, m interface{}) error {
	var err error

	// Add implementation details here

	return err
}
