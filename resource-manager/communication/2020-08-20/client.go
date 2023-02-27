package v2020_08_20

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/communication/2020-08-20/communicationservice"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Client struct {
	CommunicationService *communicationservice.CommunicationServiceClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	communicationServiceClient := communicationservice.NewCommunicationServiceClientWithBaseURI(endpoint)
	configureAuthFunc(&communicationServiceClient.Client)

	return Client{
		CommunicationService: &communicationServiceClient,
	}
}
