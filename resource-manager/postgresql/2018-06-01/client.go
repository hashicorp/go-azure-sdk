package v2018_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2018-06-01/advisors"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2018-06-01/locationbasedrecommendedactionsessionsresult"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2018-06-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2018-06-01/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2018-06-01/querytexts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2018-06-01/recommendedactions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2018-06-01/recommendedactionsessions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2018-06-01/resetqueryperformanceinsightdata"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2018-06-01/topquerystatistics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2018-06-01/waitstatistics"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Advisors                                     *advisors.AdvisorsClient
	LocationBasedRecommendedActionSessionsResult *locationbasedrecommendedactionsessionsresult.LocationBasedRecommendedActionSessionsResultClient
	PrivateEndpointConnections                   *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources                         *privatelinkresources.PrivateLinkResourcesClient
	QueryTexts                                   *querytexts.QueryTextsClient
	RecommendedActionSessions                    *recommendedactionsessions.RecommendedActionSessionsClient
	RecommendedActions                           *recommendedactions.RecommendedActionsClient
	ResetQueryPerformanceInsightData             *resetqueryperformanceinsightdata.ResetQueryPerformanceInsightDataClient
	TopQueryStatistics                           *topquerystatistics.TopQueryStatisticsClient
	WaitStatistics                               *waitstatistics.WaitStatisticsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	advisorsClient, err := advisors.NewAdvisorsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Advisors client: %+v", err)
	}
	configureFunc(advisorsClient.Client)

	locationBasedRecommendedActionSessionsResultClient, err := locationbasedrecommendedactionsessionsresult.NewLocationBasedRecommendedActionSessionsResultClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LocationBasedRecommendedActionSessionsResult client: %+v", err)
	}
	configureFunc(locationBasedRecommendedActionSessionsResultClient.Client)

	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnections client: %+v", err)
	}
	configureFunc(privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient, err := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkResources client: %+v", err)
	}
	configureFunc(privateLinkResourcesClient.Client)

	queryTextsClient, err := querytexts.NewQueryTextsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building QueryTexts client: %+v", err)
	}
	configureFunc(queryTextsClient.Client)

	recommendedActionSessionsClient, err := recommendedactionsessions.NewRecommendedActionSessionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RecommendedActionSessions client: %+v", err)
	}
	configureFunc(recommendedActionSessionsClient.Client)

	recommendedActionsClient, err := recommendedactions.NewRecommendedActionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RecommendedActions client: %+v", err)
	}
	configureFunc(recommendedActionsClient.Client)

	resetQueryPerformanceInsightDataClient, err := resetqueryperformanceinsightdata.NewResetQueryPerformanceInsightDataClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ResetQueryPerformanceInsightData client: %+v", err)
	}
	configureFunc(resetQueryPerformanceInsightDataClient.Client)

	topQueryStatisticsClient, err := topquerystatistics.NewTopQueryStatisticsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TopQueryStatistics client: %+v", err)
	}
	configureFunc(topQueryStatisticsClient.Client)

	waitStatisticsClient, err := waitstatistics.NewWaitStatisticsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WaitStatistics client: %+v", err)
	}
	configureFunc(waitStatisticsClient.Client)

	return &Client{
		Advisors: advisorsClient,
		LocationBasedRecommendedActionSessionsResult: locationBasedRecommendedActionSessionsResultClient,
		PrivateEndpointConnections:                   privateEndpointConnectionsClient,
		PrivateLinkResources:                         privateLinkResourcesClient,
		QueryTexts:                                   queryTextsClient,
		RecommendedActionSessions:                    recommendedActionSessionsClient,
		RecommendedActions:                           recommendedActionsClient,
		ResetQueryPerformanceInsightData:             resetQueryPerformanceInsightDataClient,
		TopQueryStatistics:                           topQueryStatisticsClient,
		WaitStatistics:                               waitStatisticsClient,
	}, nil
}
