package playgroundtech

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/playgroundtech/terraform-provider-playgroundtech/playgroundtech/client"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("PLAYGROUNDTECH_PASSWORD", nil),
			},
			"email": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PLAYGROUNDTECH_EMAIL", nil),
			},
		},
		ResourcesMap:   map[string]*schema.Resource{
			"playgroundtech_application": resourceApplication(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	password := d.Get("password").(string)
	email := d.Get("email").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if (password != "") && (email != "") {
		c, err := client.NewClient(nil, &password, &email)
		if err != nil {
			return nil, diag.FromErr(err)
		}

		return c, diags
	}

	c, err := client.NewClient(nil, nil, nil)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return c, diags
}