package v2024_04_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/servicelinker/2024-04-01/configurationnames"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicelinker/2024-04-01/connector"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicelinker/2024-04-01/linkers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/servicelinker/2024-04-01/servicelinker"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ConfigurationNames *configurationnames.ConfigurationNamesClient
	Connector          *connector.ConnectorClient
	Linkers            *linkers.LinkersClient
	Servicelinker      *servicelinker.ServicelinkerClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	configurationNamesClient, err := configurationnames.NewConfigurationNamesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConfigurationNames client: %+v", err)
	}
	configureFunc(configurationNamesClient.Client)

	connectorClient, err := connector.NewConnectorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Connector client: %+v", err)
	}
	configureFunc(connectorClient.Client)

	linkersClient, err := linkers.NewLinkersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Linkers client: %+v", err)
	}
	configureFunc(linkersClient.Client)

	servicelinkerClient, err := servicelinker.NewServicelinkerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Servicelinker client: %+v", err)
	}
	configureFunc(servicelinkerClient.Client)

	return &Client{
		ConfigurationNames: configurationNamesClient,
		Connector:          connectorClient,
		Linkers:            linkersClient,
		Servicelinker:      servicelinkerClient,
	}, nil
}
