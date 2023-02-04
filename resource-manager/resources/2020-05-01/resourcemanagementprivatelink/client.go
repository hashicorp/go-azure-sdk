package resourcemanagementprivatelink

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceManagementPrivateLinkClient struct {
	Client *resourcemanager.Client
}

func NewResourceManagementPrivateLinkClientWithBaseURI(api environments.Api) (*ResourceManagementPrivateLinkClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "resourcemanagementprivatelink", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ResourceManagementPrivateLinkClient: %+v", err)
	}

	return &ResourceManagementPrivateLinkClient{
		Client: client,
	}, nil
}
