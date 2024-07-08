package v2024_05_08

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/informatica/2024-05-08/organizations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/informatica/2024-05-08/serverlessruntimes"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Organizations      *organizations.OrganizationsClient
	ServerlessRuntimes *serverlessruntimes.ServerlessRuntimesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	organizationsClient, err := organizations.NewOrganizationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Organizations client: %+v", err)
	}
	configureFunc(organizationsClient.Client)

	serverlessRuntimesClient, err := serverlessruntimes.NewServerlessRuntimesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServerlessRuntimes client: %+v", err)
	}
	configureFunc(serverlessRuntimesClient.Client)

	return &Client{
		Organizations:      organizationsClient,
		ServerlessRuntimes: serverlessRuntimesClient,
	}, nil
}
