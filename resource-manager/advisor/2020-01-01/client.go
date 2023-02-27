// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2020_01_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/advisor/2020-01-01/configurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/advisor/2020-01-01/generaterecommendations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/advisor/2020-01-01/getrecommendations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/advisor/2020-01-01/metadata"
	"github.com/hashicorp/go-azure-sdk/resource-manager/advisor/2020-01-01/suppressions"
)

type Client struct {
	Configurations          *configurations.ConfigurationsClient
	GenerateRecommendations *generaterecommendations.GenerateRecommendationsClient
	GetRecommendations      *getrecommendations.GetRecommendationsClient
	Metadata                *metadata.MetadataClient
	Suppressions            *suppressions.SuppressionsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	configurationsClient := configurations.NewConfigurationsClientWithBaseURI(endpoint)
	configureAuthFunc(&configurationsClient.Client)

	generateRecommendationsClient := generaterecommendations.NewGenerateRecommendationsClientWithBaseURI(endpoint)
	configureAuthFunc(&generateRecommendationsClient.Client)

	getRecommendationsClient := getrecommendations.NewGetRecommendationsClientWithBaseURI(endpoint)
	configureAuthFunc(&getRecommendationsClient.Client)

	metadataClient := metadata.NewMetadataClientWithBaseURI(endpoint)
	configureAuthFunc(&metadataClient.Client)

	suppressionsClient := suppressions.NewSuppressionsClientWithBaseURI(endpoint)
	configureAuthFunc(&suppressionsClient.Client)

	return Client{
		Configurations:          &configurationsClient,
		GenerateRecommendations: &generateRecommendationsClient,
		GetRecommendations:      &getRecommendationsClient,
		Metadata:                &metadataClient,
		Suppressions:            &suppressionsClient,
	}
}
