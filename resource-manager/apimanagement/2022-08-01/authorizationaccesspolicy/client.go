package authorizationaccesspolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationAccessPolicyClient struct {
	Client *resourcemanager.Client
}

func NewAuthorizationAccessPolicyClientWithBaseURI(api sdkEnv.Api) (*AuthorizationAccessPolicyClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "authorizationaccesspolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthorizationAccessPolicyClient: %+v", err)
	}

	return &AuthorizationAccessPolicyClient{
		Client: client,
	}, nil
}
