package batchmanagements

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BatchManagementsClient struct {
	Client *resourcemanager.Client
}

func NewBatchManagementsClientWithBaseURI(api environments.Api) (*BatchManagementsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "batchmanagements", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BatchManagementsClient: %+v", err)
	}

	return &BatchManagementsClient{
		Client: client,
	}, nil
}
