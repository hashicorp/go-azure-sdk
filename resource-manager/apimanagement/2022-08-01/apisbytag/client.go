package apisbytag

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApisByTagClient struct {
	Client *resourcemanager.Client
}

func NewApisByTagClientWithBaseURI(api environments.Api) (*ApisByTagClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "apisbytag", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApisByTagClient: %+v", err)
	}

	return &ApisByTagClient{
		Client: client,
	}, nil
}
