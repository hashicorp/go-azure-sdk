package query

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QueryClient struct {
	Client *resourcemanager.Client
}

func NewQueryClientWithBaseURI(api sdkEnv.Api) (*QueryClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "query", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating QueryClient: %+v", err)
	}

	return &QueryClient{
		Client: client,
	}, nil
}
