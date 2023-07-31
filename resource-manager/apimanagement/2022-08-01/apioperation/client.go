package apioperation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiOperationClient struct {
	Client *resourcemanager.Client
}

func NewApiOperationClientWithBaseURI(api environments.Api) (*ApiOperationClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "apioperation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApiOperationClient: %+v", err)
	}

	return &ApiOperationClient{
		Client: client,
	}, nil
}
