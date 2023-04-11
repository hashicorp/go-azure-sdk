package v2022_05_13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/confidentialledger/2022-05-13/confidentialledger"
	"github.com/hashicorp/go-azure-sdk/resource-manager/confidentialledger/2022-05-13/nameavailability"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ConfidentialLedger *confidentialledger.ConfidentialLedgerClient
	NameAvailability   *nameavailability.NameAvailabilityClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	confidentialLedgerClient, err := confidentialledger.NewConfidentialLedgerClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ConfidentialLedger client: %+v", err)
	}
	configureFunc(confidentialLedgerClient.Client)

	nameAvailabilityClient, err := nameavailability.NewNameAvailabilityClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building NameAvailability client: %+v", err)
	}
	configureFunc(nameAvailabilityClient.Client)

	return &Client{
		ConfidentialLedger: confidentialLedgerClient,
		NameAvailability:   nameAvailabilityClient,
	}, nil
}
