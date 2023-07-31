package global

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GlobalClient struct {
	Client *resourcemanager.Client
}

func NewGlobalClientWithBaseURI(api sdkEnv.Api) (*GlobalClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "global", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GlobalClient: %+v", err)
	}

	return &GlobalClient{
		Client: client,
	}, nil
}
