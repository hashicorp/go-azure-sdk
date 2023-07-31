package adminkeys

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdminKeysClient struct {
	Client *resourcemanager.Client
}

func NewAdminKeysClientWithBaseURI(api sdkEnv.Api) (*AdminKeysClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "adminkeys", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AdminKeysClient: %+v", err)
	}

	return &AdminKeysClient{
		Client: client,
	}, nil
}
