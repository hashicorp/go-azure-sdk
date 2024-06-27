package billingaccount

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingAccountClient struct {
	Client *resourcemanager.Client
}

func NewBillingAccountClientWithBaseURI(sdkApi sdkEnv.Api) (*BillingAccountClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "billingaccount", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BillingAccountClient: %+v", err)
	}

	return &BillingAccountClient{
		Client: client,
	}, nil
}
