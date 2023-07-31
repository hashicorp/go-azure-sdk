package serverconnectionpolicies

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerConnectionPoliciesClient struct {
	Client *resourcemanager.Client
}

func NewServerConnectionPoliciesClientWithBaseURI(api sdkEnv.Api) (*ServerConnectionPoliciesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "serverconnectionpolicies", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServerConnectionPoliciesClient: %+v", err)
	}

	return &ServerConnectionPoliciesClient{
		Client: client,
	}, nil
}
