package api

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiClient struct {
	Client *resourcemanager.Client
}

func NewApiClientWithBaseURI(api environments.Api) (*ApiClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "api", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApiClient: %+v", err)
	}

	return &ApiClient{
		Client: client,
	}, nil
}
