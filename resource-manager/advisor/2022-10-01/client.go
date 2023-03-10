package v2022_10_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/advisor/2022-10-01/advisorscore"
	"github.com/hashicorp/go-azure-sdk/resource-manager/advisor/2022-10-01/configurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/advisor/2022-10-01/generaterecommendations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/advisor/2022-10-01/getrecommendations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/advisor/2022-10-01/metadata"
	"github.com/hashicorp/go-azure-sdk/resource-manager/advisor/2022-10-01/prediction"
	"github.com/hashicorp/go-azure-sdk/resource-manager/advisor/2022-10-01/suppressions"
)

type Client struct {
	AdvisorScore            *advisorscore.AdvisorScoreClient
	Configurations          *configurations.ConfigurationsClient
	GenerateRecommendations *generaterecommendations.GenerateRecommendationsClient
	GetRecommendations      *getrecommendations.GetRecommendationsClient
	Metadata                *metadata.MetadataClient
	Prediction              *prediction.PredictionClient
	Suppressions            *suppressions.SuppressionsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	advisorScoreClient := advisorscore.NewAdvisorScoreClientWithBaseURI(endpoint)
	configureAuthFunc(&advisorScoreClient.Client)

	configurationsClient := configurations.NewConfigurationsClientWithBaseURI(endpoint)
	configureAuthFunc(&configurationsClient.Client)

	generateRecommendationsClient := generaterecommendations.NewGenerateRecommendationsClientWithBaseURI(endpoint)
	configureAuthFunc(&generateRecommendationsClient.Client)

	getRecommendationsClient := getrecommendations.NewGetRecommendationsClientWithBaseURI(endpoint)
	configureAuthFunc(&getRecommendationsClient.Client)

	metadataClient := metadata.NewMetadataClientWithBaseURI(endpoint)
	configureAuthFunc(&metadataClient.Client)

	predictionClient := prediction.NewPredictionClientWithBaseURI(endpoint)
	configureAuthFunc(&predictionClient.Client)

	suppressionsClient := suppressions.NewSuppressionsClientWithBaseURI(endpoint)
	configureAuthFunc(&suppressionsClient.Client)

	return Client{
		AdvisorScore:            &advisorScoreClient,
		Configurations:          &configurationsClient,
		GenerateRecommendations: &generateRecommendationsClient,
		GetRecommendations:      &getRecommendationsClient,
		Metadata:                &metadataClient,
		Prediction:              &predictionClient,
		Suppressions:            &suppressionsClient,
	}
}
