package privatelink

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateLinkClient struct {
	Client *resourcemanager.Client
}

func NewPrivateLinkClientWithBaseURI(api sdkEnv.Api) (*PrivateLinkClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "privatelink", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivateLinkClient: %+v", err)
	}

	return &PrivateLinkClient{
		Client: client,
	}, nil
}
