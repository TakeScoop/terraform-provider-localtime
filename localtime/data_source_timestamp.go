package localtime

import (
	"context"
	"fmt"
	"time"

	"github.com/araddon/dateparse"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTimestamp() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTimestampRead,
		Schema: map[string]*schema.Schema{
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
					if err != nil {
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
					if err != nil {
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

func dataSourceTimestampRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	localTime := d.Get("local_time").(string)

	layout := d.Get("layout").(string)
	if layout == "" {
		layout_local, err := dateparse.ParseFormat(localTime)
		if err != nil {
			return diag.Errorf("error parsing local time: %s", err)
		}

		layout = fmt.Sprintf("%s %s", layout_local, d.Get("layout_timezone"))
	}

	location := time.Local
	if v, ok := d.GetOk("location"); ok {
		var err error
		location, err = time.LoadLocation(v.(string))
		if err != nil {
			return diag.Errorf("error loading location: %s", err)
		}
	}

	t, err := dateparse.ParseIn(localTime, location)
	if err != nil {
		return diag.Errorf("error parsing time in location: %s", err)
	}

	if err = d.Set("timestamp", t.Format(layout)); err != nil {
		return diag.Errorf("error setting timestamp in resource: %s", err)
	}

	d.SetId(localTime)

	return diags
}
