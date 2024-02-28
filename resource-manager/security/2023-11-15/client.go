package v2023_11_15

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2023-11-15/apimconfig"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2023-11-15/d4apicollection"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2023-11-15/d4apicollectionlist"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2023-11-15/offboardfromd4api"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2023-11-15/onboardtod4api"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	APIMConfig          *apimconfig.APIMConfigClient
	D4APICollection     *d4apicollection.D4APICollectionClient
	D4APICollectionList *d4apicollectionlist.D4APICollectionListClient
	OffboardFromD4API   *offboardfromd4api.OffboardFromD4APIClient
	OnboardToD4API      *onboardtod4api.OnboardToD4APIClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	aPIMConfigClient, err := apimconfig.NewAPIMConfigClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building APIMConfig client: %+v", err)
	}
	configureFunc(aPIMConfigClient.Client)

	d4APICollectionClient, err := d4apicollection.NewD4APICollectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building D4APICollection client: %+v", err)
	}
	configureFunc(d4APICollectionClient.Client)

	d4APICollectionListClient, err := d4apicollectionlist.NewD4APICollectionListClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building D4APICollectionList client: %+v", err)
	}
	configureFunc(d4APICollectionListClient.Client)

	offboardFromD4APIClient, err := offboardfromd4api.NewOffboardFromD4APIClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OffboardFromD4API client: %+v", err)
	}
	configureFunc(offboardFromD4APIClient.Client)

	onboardToD4APIClient, err := onboardtod4api.NewOnboardToD4APIClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnboardToD4API client: %+v", err)
	}
	configureFunc(onboardToD4APIClient.Client)

	return &Client{
		APIMConfig:          aPIMConfigClient,
		D4APICollection:     d4APICollectionClient,
		D4APICollectionList: d4APICollectionListClient,
		OffboardFromD4API:   offboardFromD4APIClient,
		OnboardToD4API:      onboardToD4APIClient,
	}, nil
}
