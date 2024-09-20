package dataproducts

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataProductsClient struct {
	Client *resourcemanager.Client
}

func NewDataProductsClientWithBaseURI(sdkApi sdkEnv.Api) (*DataProductsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "dataproducts", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DataProductsClient: %+v", err)
	}

	return &DataProductsClient{
		Client: client,
	}, nil
}
