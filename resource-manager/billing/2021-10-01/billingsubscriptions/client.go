package billingsubscriptions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingSubscriptionsClient struct {
	Client *resourcemanager.Client
}

func NewBillingSubscriptionsClientWithBaseURI(sdkApi sdkEnv.Api) (*BillingSubscriptionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "billingsubscriptions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BillingSubscriptionsClient: %+v", err)
	}

	return &BillingSubscriptionsClient{
		Client: client,
	}, nil
}
