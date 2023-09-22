package policysetdefinitionversions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicySetDefinitionVersionsClient struct {
	Client *resourcemanager.Client
}

func NewPolicySetDefinitionVersionsClientWithBaseURI(sdkApi sdkEnv.Api) (*PolicySetDefinitionVersionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "policysetdefinitionversions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicySetDefinitionVersionsClient: %+v", err)
	}

	return &PolicySetDefinitionVersionsClient{
		Client: client,
	}, nil
}
