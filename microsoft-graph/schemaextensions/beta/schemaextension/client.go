package schemaextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SchemaExtensionClient struct {
	Client *msgraph.Client
}

func NewSchemaExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*SchemaExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "schemaextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SchemaExtensionClient: %+v", err)
	}

	return &SchemaExtensionClient{
		Client: client,
	}, nil
}
