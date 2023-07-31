package elasticversions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ElasticVersionsClient struct {
	Client *resourcemanager.Client
}

func NewElasticVersionsClientWithBaseURI(api sdkEnv.Api) (*ElasticVersionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "elasticversions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ElasticVersionsClient: %+v", err)
	}

	return &ElasticVersionsClient{
		Client: client,
	}, nil
}
