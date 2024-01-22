package billingsubscriptionsaliases

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingSubscriptionsAliasesClient struct {
	Client *resourcemanager.Client
}

func NewBillingSubscriptionsAliasesClientWithBaseURI(sdkApi sdkEnv.Api) (*BillingSubscriptionsAliasesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "billingsubscriptionsaliases", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BillingSubscriptionsAliasesClient: %+v", err)
	}

	return &BillingSubscriptionsAliasesClient{
		Client: client,
	}, nil
}
