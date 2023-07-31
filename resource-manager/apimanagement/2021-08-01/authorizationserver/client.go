package authorizationserver

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationServerClient struct {
	Client *resourcemanager.Client
}

func NewAuthorizationServerClientWithBaseURI(api environments.Api) (*AuthorizationServerClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "authorizationserver", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthorizationServerClient: %+v", err)
	}

	return &AuthorizationServerClient{
		Client: client,
	}, nil
}
