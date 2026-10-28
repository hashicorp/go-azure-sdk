package origingroups

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OriginGroupsClient struct {
	Client *resourcemanager.Client
}

func NewOriginGroupsClientWithBaseURI(sdkApi sdkEnv.Api) (*OriginGroupsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "origingroups", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OriginGroupsClient: %+v", err)
	}

	return &OriginGroupsClient{
		Client: client,
	}, nil
}
