package slices

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SlicesClient struct {
	Client *resourcemanager.Client
}

func NewSlicesClientWithBaseURI(api environments.Api) (*SlicesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "slices", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SlicesClient: %+v", err)
	}

	return &SlicesClient{
		Client: client,
	}, nil
}
