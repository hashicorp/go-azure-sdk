package restore

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestoreClient struct {
	Client *resourcemanager.Client
}

func NewRestoreClientWithBaseURI(api environments.Api) (*RestoreClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "restore", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RestoreClient: %+v", err)
	}

	return &RestoreClient{
		Client: client,
	}, nil
}
