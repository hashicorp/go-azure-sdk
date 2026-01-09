package stable

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applicationtemplates/stable/applicationtemplate"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ApplicationTemplate *applicationtemplate.ApplicationTemplateClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	applicationTemplateClient, err := applicationtemplate.NewApplicationTemplateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationTemplate client: %+v", err)
	}
	configureFunc(applicationTemplateClient.Client)

	return &Client{
		ApplicationTemplate: applicationTemplateClient,
	}, nil
}
