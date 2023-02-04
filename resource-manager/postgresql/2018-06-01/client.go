package v2018_06_01

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
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
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

func NewClientWithBaseURI(api environments.Api, configureAuthFunc func(c *resourcemanager.Client)) (*Client, error) {
	advisorsClient, err := advisors.NewAdvisorsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for Advisors: %+v", err)
	}
	configureAuthFunc(advisorsClient.Client)

	locationBasedRecommendedActionSessionsResultClient, err := locationbasedrecommendedactionsessionsresult.NewLocationBasedRecommendedActionSessionsResultClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for LocationBasedRecommendedActionSessionsResult: %+v", err)
	}
	configureAuthFunc(locationBasedRecommendedActionSessionsResultClient.Client)

	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for PrivateEndpointConnections: %+v", err)
	}
	configureAuthFunc(privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient, err := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for PrivateLinkResources: %+v", err)
	}
	configureAuthFunc(privateLinkResourcesClient.Client)

	queryTextsClient, err := querytexts.NewQueryTextsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for QueryTexts: %+v", err)
	}
	configureAuthFunc(queryTextsClient.Client)

	recommendedActionSessionsClient, err := recommendedactionsessions.NewRecommendedActionSessionsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for RecommendedActionSessions: %+v", err)
	}
	configureAuthFunc(recommendedActionSessionsClient.Client)

	recommendedActionsClient, err := recommendedactions.NewRecommendedActionsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for RecommendedActions: %+v", err)
	}
	configureAuthFunc(recommendedActionsClient.Client)

	resetQueryPerformanceInsightDataClient, err := resetqueryperformanceinsightdata.NewResetQueryPerformanceInsightDataClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for ResetQueryPerformanceInsightData: %+v", err)
	}
	configureAuthFunc(resetQueryPerformanceInsightDataClient.Client)

	topQueryStatisticsClient, err := topquerystatistics.NewTopQueryStatisticsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for TopQueryStatistics: %+v", err)
	}
	configureAuthFunc(topQueryStatisticsClient.Client)

	waitStatisticsClient, err := waitstatistics.NewWaitStatisticsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for WaitStatistics: %+v", err)
	}
	configureAuthFunc(waitStatisticsClient.Client)

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
