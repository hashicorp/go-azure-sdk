package key

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeyClient struct {
	Client *resourcemanager.Client
}

func NewKeyClientWithBaseURI(api environments.Api) (*KeyClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "key", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating KeyClient: %+v", err)
	}

	return &KeyClient{
		Client: client,
	}, nil
}
