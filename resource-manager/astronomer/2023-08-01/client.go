package v2023_08_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/astronomer/2023-08-01/organizations"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Organizations *organizations.OrganizationsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	organizationsClient, err := organizations.NewOrganizationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Organizations client: %+v", err)
	}
	configureFunc(organizationsClient.Client)

	return &Client{
		Organizations: organizationsClient,
	}, nil
}
