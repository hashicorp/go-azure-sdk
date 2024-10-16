package schemaversions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SchemaVersionsClient struct {
	Client *resourcemanager.Client
}

func NewSchemaVersionsClientWithBaseURI(sdkApi sdkEnv.Api) (*SchemaVersionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "schemaversions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SchemaVersionsClient: %+v", err)
	}

	return &SchemaVersionsClient{
		Client: client,
	}, nil
}
