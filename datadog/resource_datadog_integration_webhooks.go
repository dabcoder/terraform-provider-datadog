package datadog

import (
	"fmt"
	"sync"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/zorkian/go-datadog-api"
)

var integrationWhMutex = sync.Mutex{}

func resourceDatadogIntegrationWebhooks() *schema.Resource {
	return &schema.Resource{
		Create: resourceDatadogIntegrationWebhooksCreate,
		Read:   resourceDatadogIntegrationWebhooksRead,
		Exists: resourceDatadogIntegrationWebhooksExists,
		Update: resourceDatadogIntegrationWebhooksUpdate,
		Delete: resourceDatadogIntegrationWebhooksDelete,
		Importer: &schema.ResourceImporter{
			State: resourceDatadogIntegrationWebhooksImport,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"use_custom_payload": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"custom_payload": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"encode_as_form": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"headers": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceDatadogIntegrationWebhooksCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*datadog.Client)
	integrationWhMutex.Lock()
	defer integrationWhMutex.Unlock()

	// TODO

}

func resourceDatadogIntegrationWebhooksRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*datadog.Client)

	wh, err := client.GetIntegrationWebhook()
	if err != nil {
		return err
	}

	webhooks := []map[string]string{}
	for _, webhook := range wh.Webhooks {
		webhooks = append(webhooks, map[string]string{
			"name": webhook.GetName(),
			"url":  webhook.GetURL(),
		})
	}

	d.Set("hooks", webhooks)
	d.Set("name", wh.GetName())
	d.Set("use_custom_payload", wh.HasUseCustomPayload())
	d.Set("custom_payload", wh.GetUseCustomPayload())
	d.Set("encode_as_form", wh.HasEncodeAsForm())
	d.Set("headers", wh.GetHeaders())

	return nil
}

func resourceDatadogIntegrationWebhooksDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*datadog.Client)
	integrationWhMutex.Lock()
	defer integrationWhMutex.Unlock()

	if err := client.DeleteIntegrationWebhook(); err != nil {
		return fmt.Errorf("Error while deleting a Webhooks integration: %v", err)
	}

	return nil
}
