package customer

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomerClient struct {
	Client *resourcemanager.Client
}

func NewCustomerClientWithBaseURI(sdkApi sdkEnv.Api) (*CustomerClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "customer", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CustomerClient: %+v", err)
	}

	return &CustomerClient{
		Client: client,
	}, nil
}
