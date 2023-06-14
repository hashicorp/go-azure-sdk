package elasticpooloperations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ElasticPoolOperationsClient struct {
	Client *resourcemanager.Client
}

func NewElasticPoolOperationsClientWithBaseURI(api environments.Api) (*ElasticPoolOperationsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "elasticpooloperations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ElasticPoolOperationsClient: %+v", err)
	}

	return &ElasticPoolOperationsClient{
		Client: client,
	}, nil
}
