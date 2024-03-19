package perimeterassociationproxies

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PerimeterAssociationProxiesClient struct {
	Client *resourcemanager.Client
}

func NewPerimeterAssociationProxiesClientWithBaseURI(sdkApi sdkEnv.Api) (*PerimeterAssociationProxiesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "perimeterassociationproxies", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PerimeterAssociationProxiesClient: %+v", err)
	}

	return &PerimeterAssociationProxiesClient{
		Client: client,
	}, nil
}
