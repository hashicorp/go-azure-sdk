package marketplaceregistrationdefinitions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MarketplaceRegistrationDefinitionsClient struct {
	Client *resourcemanager.Client
}

func NewMarketplaceRegistrationDefinitionsClientWithBaseURI(sdkApi sdkEnv.Api) (*MarketplaceRegistrationDefinitionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "marketplaceregistrationdefinitions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MarketplaceRegistrationDefinitionsClient: %+v", err)
	}

	return &MarketplaceRegistrationDefinitionsClient{
		Client: client,
	}, nil
}
