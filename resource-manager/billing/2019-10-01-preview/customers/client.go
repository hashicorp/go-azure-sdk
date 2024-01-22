package customers

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomersClient struct {
	Client *resourcemanager.Client
}

func NewCustomersClientWithBaseURI(sdkApi sdkEnv.Api) (*CustomersClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "customers", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CustomersClient: %+v", err)
	}

	return &CustomersClient{
		Client: client,
	}, nil
}
