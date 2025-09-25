package v2024_07_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/confluent/2024-07-01/confluents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/confluent/2024-07-01/connectorresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/confluent/2024-07-01/organizationresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/confluent/2024-07-01/scclusterrecords"
	"github.com/hashicorp/go-azure-sdk/resource-manager/confluent/2024-07-01/scenvironmentrecords"
	"github.com/hashicorp/go-azure-sdk/resource-manager/confluent/2024-07-01/topicrecords"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Confluents            *confluents.ConfluentsClient
	ConnectorResources    *connectorresources.ConnectorResourcesClient
	OrganizationResources *organizationresources.OrganizationResourcesClient
	SCClusterRecords      *scclusterrecords.SCClusterRecordsClient
	SCEnvironmentRecords  *scenvironmentrecords.SCEnvironmentRecordsClient
	TopicRecords          *topicrecords.TopicRecordsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	confluentsClient, err := confluents.NewConfluentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Confluents client: %+v", err)
	}
	configureFunc(confluentsClient.Client)

	connectorResourcesClient, err := connectorresources.NewConnectorResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConnectorResources client: %+v", err)
	}
	configureFunc(connectorResourcesClient.Client)

	organizationResourcesClient, err := organizationresources.NewOrganizationResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OrganizationResources client: %+v", err)
	}
	configureFunc(organizationResourcesClient.Client)

	sCClusterRecordsClient, err := scclusterrecords.NewSCClusterRecordsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SCClusterRecords client: %+v", err)
	}
	configureFunc(sCClusterRecordsClient.Client)

	sCEnvironmentRecordsClient, err := scenvironmentrecords.NewSCEnvironmentRecordsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SCEnvironmentRecords client: %+v", err)
	}
	configureFunc(sCEnvironmentRecordsClient.Client)

	topicRecordsClient, err := topicrecords.NewTopicRecordsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TopicRecords client: %+v", err)
	}
	configureFunc(topicRecordsClient.Client)

	return &Client{
		Confluents:            confluentsClient,
		ConnectorResources:    connectorResourcesClient,
		OrganizationResources: organizationResourcesClient,
		SCClusterRecords:      sCClusterRecordsClient,
		SCEnvironmentRecords:  sCEnvironmentRecordsClient,
		TopicRecords:          topicRecordsClient,
	}, nil
}
