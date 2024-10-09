package v2022_05_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/servicelinker/2022-05-01/links"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicelinker/2022-05-01/servicelinker"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Links         *links.LinksClient
	Servicelinker *servicelinker.ServicelinkerClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	linksClient, err := links.NewLinksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Links client: %+v", err)
	}
	configureFunc(linksClient.Client)

	servicelinkerClient, err := servicelinker.NewServicelinkerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Servicelinker client: %+v", err)
	}
	configureFunc(servicelinkerClient.Client)

	return &Client{
		Links:         linksClient,
		Servicelinker: servicelinkerClient,
	}, nil
}
