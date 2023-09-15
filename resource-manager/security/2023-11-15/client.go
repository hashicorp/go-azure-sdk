package v2023_11_15

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2023-11-15/apimconfig"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2023-11-15/d4apicollection"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2023-11-15/d4apicollectionlist"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2023-11-15/offboardfromd4api"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2023-11-15/onboardtod4api"
)

type Client struct {
	APIMConfig          *apimconfig.APIMConfigClient
	D4APICollection     *d4apicollection.D4APICollectionClient
	D4APICollectionList *d4apicollectionlist.D4APICollectionListClient
	OffboardFromD4API   *offboardfromd4api.OffboardFromD4APIClient
	OnboardToD4API      *onboardtod4api.OnboardToD4APIClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	aPIMConfigClient := apimconfig.NewAPIMConfigClientWithBaseURI(endpoint)
	configureAuthFunc(&aPIMConfigClient.Client)

	d4APICollectionClient := d4apicollection.NewD4APICollectionClientWithBaseURI(endpoint)
	configureAuthFunc(&d4APICollectionClient.Client)

	d4APICollectionListClient := d4apicollectionlist.NewD4APICollectionListClientWithBaseURI(endpoint)
	configureAuthFunc(&d4APICollectionListClient.Client)

	offboardFromD4APIClient := offboardfromd4api.NewOffboardFromD4APIClientWithBaseURI(endpoint)
	configureAuthFunc(&offboardFromD4APIClient.Client)

	onboardToD4APIClient := onboardtod4api.NewOnboardToD4APIClientWithBaseURI(endpoint)
	configureAuthFunc(&onboardToD4APIClient.Client)

	return Client{
		APIMConfig:          &aPIMConfigClient,
		D4APICollection:     &d4APICollectionClient,
		D4APICollectionList: &d4APICollectionListClient,
		OffboardFromD4API:   &offboardFromD4APIClient,
		OnboardToD4API:      &onboardToD4APIClient,
	}
}
