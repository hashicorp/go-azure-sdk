package v2021_12_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/purview/2021-12-01/account"
	"github.com/hashicorp/go-azure-sdk/resource-manager/purview/2021-12-01/defaultaccount"
	"github.com/hashicorp/go-azure-sdk/resource-manager/purview/2021-12-01/feature"
	"github.com/hashicorp/go-azure-sdk/resource-manager/purview/2021-12-01/kafkaconfiguration"
	"github.com/hashicorp/go-azure-sdk/resource-manager/purview/2021-12-01/privateendpointconnection"
	"github.com/hashicorp/go-azure-sdk/resource-manager/purview/2021-12-01/privatelinkresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/purview/2021-12-01/provider"
	"github.com/hashicorp/go-azure-sdk/resource-manager/purview/2021-12-01/usages"
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

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	accountClient := account.NewAccountClientWithBaseURI(endpoint)
	configureAuthFunc(&accountClient.Client)

	defaultAccountClient := defaultaccount.NewDefaultAccountClientWithBaseURI(endpoint)
	configureAuthFunc(&defaultAccountClient.Client)

	featureClient := feature.NewFeatureClientWithBaseURI(endpoint)
	configureAuthFunc(&featureClient.Client)

	kafkaConfigurationClient := kafkaconfiguration.NewKafkaConfigurationClientWithBaseURI(endpoint)
	configureAuthFunc(&kafkaConfigurationClient.Client)

	privateEndpointConnectionClient := privateendpointconnection.NewPrivateEndpointConnectionClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionClient.Client)

	privateLinkResourceClient := privatelinkresource.NewPrivateLinkResourceClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkResourceClient.Client)

	providerClient := provider.NewProviderClientWithBaseURI(endpoint)
	configureAuthFunc(&providerClient.Client)

	usagesClient := usages.NewUsagesClientWithBaseURI(endpoint)
	configureAuthFunc(&usagesClient.Client)

	return Client{
		Account:                   &accountClient,
		DefaultAccount:            &defaultAccountClient,
		Feature:                   &featureClient,
		KafkaConfiguration:        &kafkaConfigurationClient,
		PrivateEndpointConnection: &privateEndpointConnectionClient,
		PrivateLinkResource:       &privateLinkResourceClient,
		Provider:                  &providerClient,
		Usages:                    &usagesClient,
	}
}
