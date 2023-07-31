package graphqlapiresolver

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GraphQLApiResolverClient struct {
	Client *resourcemanager.Client
}

func NewGraphQLApiResolverClientWithBaseURI(api sdkEnv.Api) (*GraphQLApiResolverClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "graphqlapiresolver", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GraphQLApiResolverClient: %+v", err)
	}

	return &GraphQLApiResolverClient{
		Client: client,
	}, nil
}
