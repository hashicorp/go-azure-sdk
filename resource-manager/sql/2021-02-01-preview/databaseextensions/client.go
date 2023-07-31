package databaseextensions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatabaseExtensionsClient struct {
	Client *resourcemanager.Client
}

func NewDatabaseExtensionsClientWithBaseURI(api sdkEnv.Api) (*DatabaseExtensionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "databaseextensions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DatabaseExtensionsClient: %+v", err)
	}

	return &DatabaseExtensionsClient{
		Client: client,
	}, nil
}
