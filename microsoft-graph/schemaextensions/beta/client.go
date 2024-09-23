package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/schemaextensions/beta/schemaextension"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	SchemaExtension *schemaextension.SchemaExtensionClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	schemaExtensionClient, err := schemaextension.NewSchemaExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SchemaExtension client: %+v", err)
	}
	configureFunc(schemaExtensionClient.Client)

	return &Client{
		SchemaExtension: schemaExtensionClient,
	}, nil
}
