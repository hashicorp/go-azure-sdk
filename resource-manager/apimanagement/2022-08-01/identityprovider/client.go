package identityprovider

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityProviderClient struct {
	Client *resourcemanager.Client
}

func NewIdentityProviderClientWithBaseURI(api sdkEnv.Api) (*IdentityProviderClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "identityprovider", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IdentityProviderClient: %+v", err)
	}

	return &IdentityProviderClient{
		Client: client,
	}, nil
}
