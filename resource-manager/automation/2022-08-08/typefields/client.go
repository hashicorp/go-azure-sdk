package typefields

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TypeFieldsClient struct {
	Client *resourcemanager.Client
}

func NewTypeFieldsClientWithBaseURI(api environments.Api) (*TypeFieldsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "typefields", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TypeFieldsClient: %+v", err)
	}

	return &TypeFieldsClient{
		Client: client,
	}, nil
}
