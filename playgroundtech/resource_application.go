package playgroundtech

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/playgroundtech/terraform-provider-playgroundtech/playgroundtech/client"
	"strconv"
	"time"
)

func resourceApplication() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceApplicationCreate,
		ReadContext:   resourceApplicationRead,
		UpdateContext: resourceApplicationUpdate,
		DeleteContext: resourceApplicationDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"applicant_id": {
				Type: schema.TypeString,
				Computed: true,
			},
			"phone_number": {
				Type: schema.TypeString,
				Required: true,
			},
			"email": {
				Type: schema.TypeString,
				Required: true,
			},
			"linkedin": {
				Type: schema.TypeString,
				Required: true,
			},
			"github": {
				Type: schema.TypeString,
				Optional: true,
			},
			"homepage": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceApplicationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	phoneNumber := d.Get("phone_number").(string)
	email := d.Get("email").(string)
	linkedin := d.Get("linkedin").(string)
	github := d.Get("github").(string)
	homepage := d.Get("homepage").(string)

	ai := client.Application{
		PhoneNr: phoneNumber,
		Email: email,
		LinkedIn: linkedin,
		Github: github,
		HomePage: homepage,
	}

	a, err := c.CreateApplication(ai)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(a.ID))

	resourceApplicationRead(ctx, d, m)

	return diags
}

func resourceApplicationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	applicationID := d.Id()

	application, err := c.GetApplication(applicationID)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("phone_number", application.PhoneNr); err != nil {
		diag.FromErr(err)
	}
	if err := d.Set("email", application.Email); err != nil {
		diag.FromErr(err)
	}
	if err := d.Set("linkedin", application.LinkedIn); err != nil {
		diag.FromErr(err)
	}
	if err := d.Set("github", application.Github); err != nil {
		diag.FromErr(err)
	}
	if err := d.Set("homepage", application.HomePage); err != nil {
		diag.FromErr(err)
	}

	return diags
}

func resourceApplicationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)

	applicationID := d.Id()

	phoneNumber := d.Get("phone_number").(string)
	email := d.Get("email").(string)
	linkedin := d.Get("linkedin").(string)
	github := d.Get("github").(string)
	homepage := d.Get("homepage").(string)

	ai := client.Application{
		PhoneNr: phoneNumber,
		Email: email,
		LinkedIn: linkedin,
		Github: github,
		HomePage: homepage,
	}

	_, err := c.UpdateApplication(applicationID, ai)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("last_updated", time.Now().Format(time.RFC850))

	return resourceApplicationRead(ctx, d, m)
}

func resourceApplicationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.Client)
	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics

	applicationID := d.Id()

	err := c.DeleteApplication(applicationID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
