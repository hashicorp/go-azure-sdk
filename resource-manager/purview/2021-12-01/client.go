package v2021_12_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/purview/2021-12-01/account"
	"github.com/hashicorp/go-azure-sdk/resource-manager/purview/2021-12-01/defaultaccount"
	"github.com/hashicorp/go-azure-sdk/resource-manager/purview/2021-12-01/feature"
	"github.com/hashicorp/go-azure-sdk/resource-manager/purview/2021-12-01/kafkaconfiguration"
	"github.com/hashicorp/go-azure-sdk/resource-manager/purview/2021-12-01/privateendpointconnection"
	"github.com/hashicorp/go-azure-sdk/resource-manager/purview/2021-12-01/privatelinkresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/purview/2021-12-01/provider"
	"github.com/hashicorp/go-azure-sdk/resource-manager/purview/2021-12-01/usages"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Account                   *account.AccountClient
	DefaultAccount            *defaultaccount.DefaultAccountClient
	Feature                   *feature.FeatureClient
	KafkaConfiguration        *kafkaconfiguration.KafkaConfigurationClient
	PrivateEndpointConnection *privateendpointconnection.PrivateEndpointConnectionClient
	PrivateLinkResource       *privatelinkresource.PrivateLinkResourceClient
	Provider                  *provider.ProviderClient
	Usages                    *usages.UsagesClient
}

func NewClientWithBaseURI(api sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	accountClient, err := account.NewAccountClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Account client: %+v", err)
	}
	configureFunc(accountClient.Client)

	defaultAccountClient, err := defaultaccount.NewDefaultAccountClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building DefaultAccount client: %+v", err)
	}
	configureFunc(defaultAccountClient.Client)

	featureClient, err := feature.NewFeatureClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Feature client: %+v", err)
	}
	configureFunc(featureClient.Client)

	kafkaConfigurationClient, err := kafkaconfiguration.NewKafkaConfigurationClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building KafkaConfiguration client: %+v", err)
	}
	configureFunc(kafkaConfigurationClient.Client)

	privateEndpointConnectionClient, err := privateendpointconnection.NewPrivateEndpointConnectionClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnection client: %+v", err)
	}
	configureFunc(privateEndpointConnectionClient.Client)

	privateLinkResourceClient, err := privatelinkresource.NewPrivateLinkResourceClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkResource client: %+v", err)
	}
	configureFunc(privateLinkResourceClient.Client)

	providerClient, err := provider.NewProviderClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Provider client: %+v", err)
	}
	configureFunc(providerClient.Client)

	usagesClient, err := usages.NewUsagesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Usages client: %+v", err)
	}
	configureFunc(usagesClient.Client)

	return &Client{
		Account:                   accountClient,
		DefaultAccount:            defaultAccountClient,
		Feature:                   featureClient,
		KafkaConfiguration:        kafkaConfigurationClient,
		PrivateEndpointConnection: privateEndpointConnectionClient,
		PrivateLinkResource:       privateLinkResourceClient,
		Provider:                  providerClient,
		Usages:                    usagesClient,
	}, nil
}
