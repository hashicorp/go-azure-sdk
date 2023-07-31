package deletedservice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletedServiceClient struct {
	Client *resourcemanager.Client
}

func NewDeletedServiceClientWithBaseURI(api environments.Api) (*DeletedServiceClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "deletedservice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeletedServiceClient: %+v", err)
	}

	return &DeletedServiceClient{
		Client: client,
	}, nil
}
