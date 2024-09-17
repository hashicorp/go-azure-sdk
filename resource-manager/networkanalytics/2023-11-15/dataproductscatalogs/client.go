package dataproductscatalogs

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataProductsCatalogsClient struct {
	Client *resourcemanager.Client
}

func NewDataProductsCatalogsClientWithBaseURI(sdkApi sdkEnv.Api) (*DataProductsCatalogsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "dataproductscatalogs", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DataProductsCatalogsClient: %+v", err)
	}

	return &DataProductsCatalogsClient{
		Client: client,
	}, nil
}
