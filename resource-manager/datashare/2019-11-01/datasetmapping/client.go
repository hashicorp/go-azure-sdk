package datasetmapping

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataSetMappingClient struct {
	Client *resourcemanager.Client
}

func NewDataSetMappingClientWithBaseURI(api environments.Api) (*DataSetMappingClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "datasetmapping", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DataSetMappingClient: %+v", err)
	}

	return &DataSetMappingClient{
		Client: client,
	}, nil
}
