// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2018_06_01

import (
	"github.com/Azure/go-autorest/autorest"
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

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	advisorsClient := advisors.NewAdvisorsClientWithBaseURI(endpoint)
	configureAuthFunc(&advisorsClient.Client)

	locationBasedRecommendedActionSessionsResultClient := locationbasedrecommendedactionsessionsresult.NewLocationBasedRecommendedActionSessionsResultClientWithBaseURI(endpoint)
	configureAuthFunc(&locationBasedRecommendedActionSessionsResultClient.Client)

	privateEndpointConnectionsClient := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkResourcesClient.Client)

	queryTextsClient := querytexts.NewQueryTextsClientWithBaseURI(endpoint)
	configureAuthFunc(&queryTextsClient.Client)

	recommendedActionSessionsClient := recommendedactionsessions.NewRecommendedActionSessionsClientWithBaseURI(endpoint)
	configureAuthFunc(&recommendedActionSessionsClient.Client)

	recommendedActionsClient := recommendedactions.NewRecommendedActionsClientWithBaseURI(endpoint)
	configureAuthFunc(&recommendedActionsClient.Client)

	resetQueryPerformanceInsightDataClient := resetqueryperformanceinsightdata.NewResetQueryPerformanceInsightDataClientWithBaseURI(endpoint)
	configureAuthFunc(&resetQueryPerformanceInsightDataClient.Client)

	topQueryStatisticsClient := topquerystatistics.NewTopQueryStatisticsClientWithBaseURI(endpoint)
	configureAuthFunc(&topQueryStatisticsClient.Client)

	waitStatisticsClient := waitstatistics.NewWaitStatisticsClientWithBaseURI(endpoint)
	configureAuthFunc(&waitStatisticsClient.Client)

	return Client{
		Advisors: &advisorsClient,
		LocationBasedRecommendedActionSessionsResult: &locationBasedRecommendedActionSessionsResultClient,
		PrivateEndpointConnections:                   &privateEndpointConnectionsClient,
		PrivateLinkResources:                         &privateLinkResourcesClient,
		QueryTexts:                                   &queryTextsClient,
		RecommendedActionSessions:                    &recommendedActionSessionsClient,
		RecommendedActions:                           &recommendedActionsClient,
		ResetQueryPerformanceInsightData:             &resetQueryPerformanceInsightDataClient,
		TopQueryStatistics:                           &topQueryStatisticsClient,
		WaitStatistics:                               &waitStatisticsClient,
	}
}
