package authorizationloginlinks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationLoginLinksClient struct {
	Client *resourcemanager.Client
}

func NewAuthorizationLoginLinksClientWithBaseURI(api environments.Api) (*AuthorizationLoginLinksClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "authorizationloginlinks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthorizationLoginLinksClient: %+v", err)
	}

	return &AuthorizationLoginLinksClient{
		Client: client,
	}, nil
}
