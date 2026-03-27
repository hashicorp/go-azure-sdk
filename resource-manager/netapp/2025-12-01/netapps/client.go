package netapps

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetappsClient struct {
	Client *resourcemanager.Client
}

func NewNetappsClientWithBaseURI(sdkApi sdkEnv.Api) (*NetappsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "netapps", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating NetappsClient: %+v", err)
	}

	return &NetappsClient{
		Client: client,
	}, nil
}
