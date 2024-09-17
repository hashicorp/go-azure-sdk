package elasticpooloperations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ElasticPoolOperationsClient struct {
	Client *resourcemanager.Client
}

func NewElasticPoolOperationsClientWithBaseURI(sdkApi sdkEnv.Api) (*ElasticPoolOperationsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "elasticpooloperations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ElasticPoolOperationsClient: %+v", err)
	}

	return &ElasticPoolOperationsClient{
		Client: client,
	}, nil
}
