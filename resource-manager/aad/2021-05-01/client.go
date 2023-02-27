// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2021_05_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/aad/2021-05-01/domainservices"
	"github.com/hashicorp/go-azure-sdk/resource-manager/aad/2021-05-01/oucontainer"
)

type Client struct {
	DomainServices *domainservices.DomainServicesClient
	OuContainer    *oucontainer.OuContainerClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	domainServicesClient := domainservices.NewDomainServicesClientWithBaseURI(endpoint)
	configureAuthFunc(&domainServicesClient.Client)

	ouContainerClient := oucontainer.NewOuContainerClientWithBaseURI(endpoint)
	configureAuthFunc(&ouContainerClient.Client)

	return Client{
		DomainServices: &domainServicesClient,
		OuContainer:    &ouContainerClient,
	}
}
