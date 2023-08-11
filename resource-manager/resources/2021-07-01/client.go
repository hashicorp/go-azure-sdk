package v2021_07_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2021-07-01/features"
	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2021-07-01/subscriptionfeatureregistrations"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Features                         *features.FeaturesClient
	SubscriptionFeatureRegistrations *subscriptionfeatureregistrations.SubscriptionFeatureRegistrationsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	featuresClient, err := features.NewFeaturesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Features client: %+v", err)
	}
	configureFunc(featuresClient.Client)

	subscriptionFeatureRegistrationsClient, err := subscriptionfeatureregistrations.NewSubscriptionFeatureRegistrationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SubscriptionFeatureRegistrations client: %+v", err)
	}
	configureFunc(subscriptionFeatureRegistrationsClient.Client)

	return &Client{
		Features:                         featuresClient,
		SubscriptionFeatureRegistrations: subscriptionFeatureRegistrationsClient,
	}, nil
}
