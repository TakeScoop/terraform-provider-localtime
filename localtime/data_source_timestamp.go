package hashicups

import (
	"time"

	"github.com/araddon/dateparse"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTimestamp() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTimestampRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"layout": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"layout_timezone"},
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					t := time.Now()

					if v == t.Format(v) {
						errs = append(errs, fmt.Errorf("'%s' is not a valid time layout!", v))
					}
					return
				},
			},
			"layout_timezone": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				Default:       "-0700",
				ConflictsWith: []string{"layout"},
			},
			"local_time": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					_, err := dateparse.ParseStrict(v)
					if err {
						errs = append(errs, err)
					}
					return
				},
			},
			"location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					_, err := time.LoadLocation(v)
					if err {
						errs = append(errs, err)
					}
					return
				},
			},
			"timestamp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceTimestampRead(ctx context.Context, d *schema.ResourceData, m interface{}) error {
	var err error

	// Add implementation details here

	return err
}
