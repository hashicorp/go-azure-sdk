package v1_0

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/applicationtemplates/v1_0/applicationtemplate"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
)

type Client struct {
	ApplicationTemplate *applicationtemplate.ApplicationTemplateClient
}

func NewClientWithBaseURI(api skdEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	applicationTemplateClient, err := applicationtemplate.NewApplicationTemplateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationTemplate client: %+v", err)
	}
	configureFunc(applicationTemplateClient.Client)

	return &Client{
		ApplicationTemplate: applicationTemplateClient,
	}, nil
}
