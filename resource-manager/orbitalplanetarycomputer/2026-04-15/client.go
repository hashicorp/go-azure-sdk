package v2026_04_15

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/orbitalplanetarycomputer/2026-04-15/geocatalogs"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	GeoCatalogs *geocatalogs.GeoCatalogsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	geoCatalogsClient, err := geocatalogs.NewGeoCatalogsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GeoCatalogs client: %+v", err)
	}
	configureFunc(geoCatalogsClient.Client)

	return &Client{
		GeoCatalogs: geoCatalogsClient,
	}, nil
}
