package policydefinitions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyDefinitionsClient struct {
	Client *resourcemanager.Client
}

func NewPolicyDefinitionsClientWithBaseURI(api environments.Api) (*PolicyDefinitionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "policydefinitions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyDefinitionsClient: %+v", err)
	}

	return &PolicyDefinitionsClient{
		Client: client,
	}, nil
}
