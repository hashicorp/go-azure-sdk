package v2015_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/marketplaceordering/2015-06-01/agreements"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Agreements *agreements.AgreementsClient
}

func NewClientWithBaseURI(api sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	agreementsClient, err := agreements.NewAgreementsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Agreements client: %+v", err)
	}
	configureFunc(agreementsClient.Client)

	return &Client{
		Agreements: agreementsClient,
	}, nil
}
