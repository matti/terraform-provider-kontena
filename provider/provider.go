package provider

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/kontena/terraform-provider-kontena/client"
)

func provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"access_token": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"kontena_grid": resourceKontenaGrid(),
		},
		ConfigureFunc: providerConfigure,
	}
}

type providerMeta struct {
	client *client.Client
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	var config = client.Config{
		URL:         d.Get("url").(string),
		AccessToken: d.Get("access_token").(string),
	}

	if client, err := config.MakeClient(); err != nil {
		return nil, err
	} else {
		var meta = providerMeta{
			client: client,
		}

		log.Printf("[INFO] Kontena: client %v", meta.client)

		return meta, nil
	}
}

func Provider() terraform.ResourceProvider {
	return provider()
}