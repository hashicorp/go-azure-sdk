package policydefinitionversions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyDefinitionVersionsClient struct {
	Client *resourcemanager.Client
}

func NewPolicyDefinitionVersionsClientWithBaseURI(sdkApi sdkEnv.Api) (*PolicyDefinitionVersionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "policydefinitionversions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyDefinitionVersionsClient: %+v", err)
	}

	return &PolicyDefinitionVersionsClient{
		Client: client,
	}, nil
}
