package associatedtenant

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssociatedTenantClient struct {
	Client *resourcemanager.Client
}

func NewAssociatedTenantClientWithBaseURI(sdkApi sdkEnv.Api) (*AssociatedTenantClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "associatedtenant", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AssociatedTenantClient: %+v", err)
	}

	return &AssociatedTenantClient{
		Client: client,
	}, nil
}
