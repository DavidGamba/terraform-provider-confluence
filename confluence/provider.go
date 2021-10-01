package confluence

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider returns the ResourceProvider for Confluence
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"site": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Cloud Confluence Site Name (the name before atlassian.net)",
				DefaultFunc: schema.EnvDefaultFunc("CONFLUENCE_SITE", nil),
			},
			"user": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "User's email address",
				DefaultFunc: schema.EnvDefaultFunc("CONFLUENCE_USER", nil),
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Confluence API Token for user",
				DefaultFunc: schema.EnvDefaultFunc("CONFLUENCE_TOKEN", nil),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap: map[string]*schema.Resource{
			"confluence_content":    resourceContent(),
			"confluence_attachment": resourceAttachment(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return NewClient(&NewClientInput{
		site:  d.Get("site").(string),
		token: d.Get("token").(string),
		user:  d.Get("user").(string),
	}), nil
}
