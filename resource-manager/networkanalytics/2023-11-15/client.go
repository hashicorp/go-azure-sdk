package v2023_11_15

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/networkanalytics/2023-11-15/dataproducts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/networkanalytics/2023-11-15/dataproductscatalogs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/networkanalytics/2023-11-15/datatypes"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	DataProducts         *dataproducts.DataProductsClient
	DataProductsCatalogs *dataproductscatalogs.DataProductsCatalogsClient
	DataTypes            *datatypes.DataTypesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	dataProductsCatalogsClient, err := dataproductscatalogs.NewDataProductsCatalogsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataProductsCatalogs client: %+v", err)
	}
	configureFunc(dataProductsCatalogsClient.Client)

	dataProductsClient, err := dataproducts.NewDataProductsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataProducts client: %+v", err)
	}
	configureFunc(dataProductsClient.Client)

	dataTypesClient, err := datatypes.NewDataTypesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataTypes client: %+v", err)
	}
	configureFunc(dataTypesClient.Client)

	return &Client{
		DataProducts:         dataProductsClient,
		DataProductsCatalogs: dataProductsCatalogsClient,
		DataTypes:            dataTypesClient,
	}, nil
}
