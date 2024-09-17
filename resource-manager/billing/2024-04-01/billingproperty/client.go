package billingproperty

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingPropertyClient struct {
	Client *resourcemanager.Client
}

func NewBillingPropertyClientWithBaseURI(sdkApi sdkEnv.Api) (*BillingPropertyClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "billingproperty", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BillingPropertyClient: %+v", err)
	}

	return &BillingPropertyClient{
		Client: client,
	}, nil
}
