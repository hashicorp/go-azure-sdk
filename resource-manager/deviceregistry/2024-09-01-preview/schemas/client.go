package schemas

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SchemasClient struct {
	Client *resourcemanager.Client
}

func NewSchemasClientWithBaseURI(sdkApi sdkEnv.Api) (*SchemasClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "schemas", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SchemasClient: %+v", err)
	}

	return &SchemasClient{
		Client: client,
	}, nil
}
