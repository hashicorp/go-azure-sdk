package apioperationsbytag

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiOperationsByTagClient struct {
	Client *resourcemanager.Client
}

func NewApiOperationsByTagClientWithBaseURI(api sdkEnv.Api) (*ApiOperationsByTagClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "apioperationsbytag", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApiOperationsByTagClient: %+v", err)
	}

	return &ApiOperationsByTagClient{
		Client: client,
	}, nil
}
