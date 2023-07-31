package customapis

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomAPIsClient struct {
	Client *resourcemanager.Client
}

func NewCustomAPIsClientWithBaseURI(api sdkEnv.Api) (*CustomAPIsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "customapis", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CustomAPIsClient: %+v", err)
	}

	return &CustomAPIsClient{
		Client: client,
	}, nil
}
