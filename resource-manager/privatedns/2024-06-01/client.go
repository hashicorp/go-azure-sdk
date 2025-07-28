package v2024_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/privatedns/2024-06-01/privatedns"
	"github.com/hashicorp/go-azure-sdk/resource-manager/privatedns/2024-06-01/privatezones"
	"github.com/hashicorp/go-azure-sdk/resource-manager/privatedns/2024-06-01/virtualnetworklinks"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	PrivateDNS          *privatedns.PrivateDNSClient
	PrivateZones        *privatezones.PrivateZonesClient
	VirtualNetworkLinks *virtualnetworklinks.VirtualNetworkLinksClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	privateDNSClient, err := privatedns.NewPrivateDNSClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateDNS client: %+v", err)
	}
	configureFunc(privateDNSClient.Client)

	privateZonesClient, err := privatezones.NewPrivateZonesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateZones client: %+v", err)
	}
	configureFunc(privateZonesClient.Client)

	virtualNetworkLinksClient, err := virtualnetworklinks.NewVirtualNetworkLinksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualNetworkLinks client: %+v", err)
	}
	configureFunc(virtualNetworkLinksClient.Client)

	return &Client{
		PrivateDNS:          privateDNSClient,
		PrivateZones:        privateZonesClient,
		VirtualNetworkLinks: virtualNetworkLinksClient,
	}, nil
}
