package apipolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiPolicyClient struct {
	Client *resourcemanager.Client
}

func NewApiPolicyClientWithBaseURI(api sdkEnv.Api) (*ApiPolicyClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "apipolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApiPolicyClient: %+v", err)
	}

	return &ApiPolicyClient{
		Client: client,
	}, nil
}
