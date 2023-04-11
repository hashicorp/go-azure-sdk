package v2020_08_20

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/communication/2020-08-20/communicationservice"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	CommunicationService *communicationservice.CommunicationServiceClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	communicationServiceClient, err := communicationservice.NewCommunicationServiceClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building CommunicationService client: %+v", err)
	}
	configureFunc(communicationServiceClient.Client)

	return &Client{
		CommunicationService: communicationServiceClient,
	}, nil
}
