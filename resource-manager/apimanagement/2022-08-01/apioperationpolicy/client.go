package apioperationpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiOperationPolicyClient struct {
	Client *resourcemanager.Client
}

func NewApiOperationPolicyClientWithBaseURI(api sdkEnv.Api) (*ApiOperationPolicyClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "apioperationpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApiOperationPolicyClient: %+v", err)
	}

	return &ApiOperationPolicyClient{
		Client: client,
	}, nil
}
