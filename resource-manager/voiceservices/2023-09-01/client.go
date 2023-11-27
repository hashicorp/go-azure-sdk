package v2023_09_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/voiceservices/2023-09-01/communicationsgateways"
	"github.com/hashicorp/go-azure-sdk/resource-manager/voiceservices/2023-09-01/testlines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/voiceservices/2023-09-01/voiceservices"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	CommunicationsGateways *communicationsgateways.CommunicationsGatewaysClient
	TestLines              *testlines.TestLinesClient
	Voiceservices          *voiceservices.VoiceservicesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	communicationsGatewaysClient, err := communicationsgateways.NewCommunicationsGatewaysClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CommunicationsGateways client: %+v", err)
	}
	configureFunc(communicationsGatewaysClient.Client)

	testLinesClient, err := testlines.NewTestLinesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TestLines client: %+v", err)
	}
	configureFunc(testLinesClient.Client)

	voiceservicesClient, err := voiceservices.NewVoiceservicesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Voiceservices client: %+v", err)
	}
	configureFunc(voiceservicesClient.Client)

	return &Client{
		CommunicationsGateways: communicationsGatewaysClient,
		TestLines:              testLinesClient,
		Voiceservices:          voiceservicesClient,
	}, nil
}
