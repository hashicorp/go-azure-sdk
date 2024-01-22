package products

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProductsClient struct {
	Client *resourcemanager.Client
}

func NewProductsClientWithBaseURI(sdkApi sdkEnv.Api) (*ProductsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "products", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProductsClient: %+v", err)
	}

	return &ProductsClient{
		Client: client,
	}, nil
}
