package operations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationsClient struct {
	Client *resourcemanager.Client
}

func NewOperationsClientWithBaseURI(api environments.Api) (*OperationsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "operations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OperationsClient: %+v", err)
	}

	return &OperationsClient{
		Client: client,
	}, nil
}
