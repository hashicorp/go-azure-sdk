package billingmeters

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingMetersClient struct {
	Client *resourcemanager.Client
}

func NewBillingMetersClientWithBaseURI(sdkApi sdkEnv.Api) (*BillingMetersClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "billingmeters", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BillingMetersClient: %+v", err)
	}

	return &BillingMetersClient{
		Client: client,
	}, nil
}
