package cloudlinks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudLinksClient struct {
	Client *resourcemanager.Client
}

func NewCloudLinksClientWithBaseURI(api environments.Api) (*CloudLinksClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "cloudlinks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CloudLinksClient: %+v", err)
	}

	return &CloudLinksClient{
		Client: client,
	}, nil
}
