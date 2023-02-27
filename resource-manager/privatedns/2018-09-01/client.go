// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2018_09_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/privatedns/2018-09-01/privatezones"
	"github.com/hashicorp/go-azure-sdk/resource-manager/privatedns/2018-09-01/recordsets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/privatedns/2018-09-01/virtualnetworklinks"
)

type Client struct {
	PrivateZones        *privatezones.PrivateZonesClient
	RecordSets          *recordsets.RecordSetsClient
	VirtualNetworkLinks *virtualnetworklinks.VirtualNetworkLinksClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	privateZonesClient := privatezones.NewPrivateZonesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateZonesClient.Client)

	recordSetsClient := recordsets.NewRecordSetsClientWithBaseURI(endpoint)
	configureAuthFunc(&recordSetsClient.Client)

	virtualNetworkLinksClient := virtualnetworklinks.NewVirtualNetworkLinksClientWithBaseURI(endpoint)
	configureAuthFunc(&virtualNetworkLinksClient.Client)

	return Client{
		PrivateZones:        &privateZonesClient,
		RecordSets:          &recordSetsClient,
		VirtualNetworkLinks: &virtualNetworkLinksClient,
	}
}
