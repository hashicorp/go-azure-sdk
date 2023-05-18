package usagedetails

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UsageDetailsClient struct {
	Client *resourcemanager.Client
}

func NewUsageDetailsClientWithBaseURI(api environments.Api) (*UsageDetailsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "usagedetails", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UsageDetailsClient: %+v", err)
	}

	return &UsageDetailsClient{
		Client: client,
	}, nil
}
