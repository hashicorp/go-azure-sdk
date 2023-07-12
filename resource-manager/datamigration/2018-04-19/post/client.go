package post

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type POSTClient struct {
	Client *resourcemanager.Client
}

func NewPOSTClientWithBaseURI(api environments.Api) (*POSTClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "post", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating POSTClient: %+v", err)
	}

	return &POSTClient{
		Client: client,
	}, nil
}
