package v2023_01_31

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/voiceservices/2023-01-31/communicationsgateways"
	"github.com/hashicorp/go-azure-sdk/resource-manager/voiceservices/2023-01-31/testlines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/voiceservices/2023-01-31/voiceservices"
)

type Client struct {
	CommunicationsGateways *communicationsgateways.CommunicationsGatewaysClient
	TestLines              *testlines.TestLinesClient
	Voiceservices          *voiceservices.VoiceservicesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	communicationsGatewaysClient := communicationsgateways.NewCommunicationsGatewaysClientWithBaseURI(endpoint)
	configureAuthFunc(&communicationsGatewaysClient.Client)

	testLinesClient := testlines.NewTestLinesClientWithBaseURI(endpoint)
	configureAuthFunc(&testLinesClient.Client)

	voiceservicesClient := voiceservices.NewVoiceservicesClientWithBaseURI(endpoint)
	configureAuthFunc(&voiceservicesClient.Client)

	return Client{
		CommunicationsGateways: &communicationsGatewaysClient,
		TestLines:              &testLinesClient,
		Voiceservices:          &voiceservicesClient,
	}
}
