package v2023_03_31

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/communication/2023-03-31/communicationservices"
	"github.com/hashicorp/go-azure-sdk/resource-manager/communication/2023-03-31/domains"
	"github.com/hashicorp/go-azure-sdk/resource-manager/communication/2023-03-31/emailservices"
	"github.com/hashicorp/go-azure-sdk/resource-manager/communication/2023-03-31/senderusernames"
)

type Client struct {
	CommunicationServices *communicationservices.CommunicationServicesClient
	Domains               *domains.DomainsClient
	EmailServices         *emailservices.EmailServicesClient
	SenderUsernames       *senderusernames.SenderUsernamesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	communicationServicesClient := communicationservices.NewCommunicationServicesClientWithBaseURI(endpoint)
	configureAuthFunc(&communicationServicesClient.Client)

	domainsClient := domains.NewDomainsClientWithBaseURI(endpoint)
	configureAuthFunc(&domainsClient.Client)

	emailServicesClient := emailservices.NewEmailServicesClientWithBaseURI(endpoint)
	configureAuthFunc(&emailServicesClient.Client)

	senderUsernamesClient := senderusernames.NewSenderUsernamesClientWithBaseURI(endpoint)
	configureAuthFunc(&senderUsernamesClient.Client)

	return Client{
		CommunicationServices: &communicationServicesClient,
		Domains:               &domainsClient,
		EmailServices:         &emailServicesClient,
		SenderUsernames:       &senderUsernamesClient,
	}
}
