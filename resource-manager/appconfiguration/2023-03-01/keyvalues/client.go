package keyvalues

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeyValuesClient struct {
	Client *resourcemanager.Client
}

func NewKeyValuesClientWithBaseURI(api environments.Api) (*KeyValuesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "keyvalues", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating KeyValuesClient: %+v", err)
	}

	return &KeyValuesClient{
		Client: client,
	}, nil
}
