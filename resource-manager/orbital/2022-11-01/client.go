package v2022_11_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/orbital/2022-11-01/contact"
	"github.com/hashicorp/go-azure-sdk/resource-manager/orbital/2022-11-01/contactprofile"
	"github.com/hashicorp/go-azure-sdk/resource-manager/orbital/2022-11-01/groundstation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/orbital/2022-11-01/spacecraft"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Client struct {
	Contact        *contact.ContactClient
	ContactProfile *contactprofile.ContactProfileClient
	GroundStation  *groundstation.GroundStationClient
	Spacecraft     *spacecraft.SpacecraftClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	contactClient := contact.NewContactClientWithBaseURI(endpoint)
	configureAuthFunc(&contactClient.Client)

	contactProfileClient := contactprofile.NewContactProfileClientWithBaseURI(endpoint)
	configureAuthFunc(&contactProfileClient.Client)

	groundStationClient := groundstation.NewGroundStationClientWithBaseURI(endpoint)
	configureAuthFunc(&groundStationClient.Client)

	spacecraftClient := spacecraft.NewSpacecraftClientWithBaseURI(endpoint)
	configureAuthFunc(&spacecraftClient.Client)

	return Client{
		Contact:        &contactClient,
		ContactProfile: &contactProfileClient,
		GroundStation:  &groundStationClient,
		Spacecraft:     &spacecraftClient,
	}
}
