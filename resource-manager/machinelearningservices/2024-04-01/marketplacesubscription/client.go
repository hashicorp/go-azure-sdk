package marketplacesubscription

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MarketplaceSubscriptionClient struct {
	Client *resourcemanager.Client
}

func NewMarketplaceSubscriptionClientWithBaseURI(sdkApi sdkEnv.Api) (*MarketplaceSubscriptionClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "marketplacesubscription", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MarketplaceSubscriptionClient: %+v", err)
	}

	return &MarketplaceSubscriptionClient{
		Client: client,
	}, nil
}
