package pricesheet

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PriceSheetClient struct {
	Client *resourcemanager.Client
}

func NewPriceSheetClientWithBaseURI(sdkApi sdkEnv.Api) (*PriceSheetClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "pricesheet", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PriceSheetClient: %+v", err)
	}

	return &PriceSheetClient{
		Client: client,
	}, nil
}
