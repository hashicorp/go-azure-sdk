package v2023_01_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/advisor/2023-01-01/advisorscore"
	"github.com/hashicorp/go-azure-sdk/resource-manager/advisor/2023-01-01/configurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/advisor/2023-01-01/generaterecommendations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/advisor/2023-01-01/getrecommendations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/advisor/2023-01-01/metadata"
	"github.com/hashicorp/go-azure-sdk/resource-manager/advisor/2023-01-01/prediction"
	"github.com/hashicorp/go-azure-sdk/resource-manager/advisor/2023-01-01/suppressions"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
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

func NewClientWithBaseURI(api sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	advisorScoreClient, err := advisorscore.NewAdvisorScoreClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building AdvisorScore client: %+v", err)
	}
	configureFunc(advisorScoreClient.Client)

	configurationsClient, err := configurations.NewConfigurationsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Configurations client: %+v", err)
	}
	configureFunc(configurationsClient.Client)

	generateRecommendationsClient, err := generaterecommendations.NewGenerateRecommendationsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building GenerateRecommendations client: %+v", err)
	}
	configureFunc(generateRecommendationsClient.Client)

	getRecommendationsClient, err := getrecommendations.NewGetRecommendationsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building GetRecommendations client: %+v", err)
	}
	configureFunc(getRecommendationsClient.Client)

	metadataClient, err := metadata.NewMetadataClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Metadata client: %+v", err)
	}
	configureFunc(metadataClient.Client)

	predictionClient, err := prediction.NewPredictionClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Prediction client: %+v", err)
	}
	configureFunc(predictionClient.Client)

	suppressionsClient, err := suppressions.NewSuppressionsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Suppressions client: %+v", err)
	}
	configureFunc(suppressionsClient.Client)

	return &Client{
		AdvisorScore:            advisorScoreClient,
		Configurations:          configurationsClient,
		GenerateRecommendations: generateRecommendationsClient,
		GetRecommendations:      getRecommendationsClient,
		Metadata:                metadataClient,
		Prediction:              predictionClient,
		Suppressions:            suppressionsClient,
	}, nil
}
