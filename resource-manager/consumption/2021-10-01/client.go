// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2021_10_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2021-10-01/aggregatedcost"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2021-10-01/balances"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2021-10-01/budgets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2021-10-01/charges"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2021-10-01/credits"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2021-10-01/events"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2021-10-01/lots"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2021-10-01/marketplaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2021-10-01/pricesheet"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2021-10-01/reservationdetails"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2021-10-01/reservationrecommendationdetails"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2021-10-01/reservationrecommendations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2021-10-01/reservationsummaries"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2021-10-01/reservationtransactions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2021-10-01/usagedetails"
)

type Client struct {
	AggregatedCost                   *aggregatedcost.AggregatedCostClient
	Balances                         *balances.BalancesClient
	Budgets                          *budgets.BudgetsClient
	Charges                          *charges.ChargesClient
	Credits                          *credits.CreditsClient
	Events                           *events.EventsClient
	Lots                             *lots.LotsClient
	Marketplaces                     *marketplaces.MarketplacesClient
	PriceSheet                       *pricesheet.PriceSheetClient
	ReservationDetails               *reservationdetails.ReservationDetailsClient
	ReservationRecommendationDetails *reservationrecommendationdetails.ReservationRecommendationDetailsClient
	ReservationRecommendations       *reservationrecommendations.ReservationRecommendationsClient
	ReservationSummaries             *reservationsummaries.ReservationSummariesClient
	ReservationTransactions          *reservationtransactions.ReservationTransactionsClient
	UsageDetails                     *usagedetails.UsageDetailsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	aggregatedCostClient := aggregatedcost.NewAggregatedCostClientWithBaseURI(endpoint)
	configureAuthFunc(&aggregatedCostClient.Client)

	balancesClient := balances.NewBalancesClientWithBaseURI(endpoint)
	configureAuthFunc(&balancesClient.Client)

	budgetsClient := budgets.NewBudgetsClientWithBaseURI(endpoint)
	configureAuthFunc(&budgetsClient.Client)

	chargesClient := charges.NewChargesClientWithBaseURI(endpoint)
	configureAuthFunc(&chargesClient.Client)

	creditsClient := credits.NewCreditsClientWithBaseURI(endpoint)
	configureAuthFunc(&creditsClient.Client)

	eventsClient := events.NewEventsClientWithBaseURI(endpoint)
	configureAuthFunc(&eventsClient.Client)

	lotsClient := lots.NewLotsClientWithBaseURI(endpoint)
	configureAuthFunc(&lotsClient.Client)

	marketplacesClient := marketplaces.NewMarketplacesClientWithBaseURI(endpoint)
	configureAuthFunc(&marketplacesClient.Client)

	priceSheetClient := pricesheet.NewPriceSheetClientWithBaseURI(endpoint)
	configureAuthFunc(&priceSheetClient.Client)

	reservationDetailsClient := reservationdetails.NewReservationDetailsClientWithBaseURI(endpoint)
	configureAuthFunc(&reservationDetailsClient.Client)

	reservationRecommendationDetailsClient := reservationrecommendationdetails.NewReservationRecommendationDetailsClientWithBaseURI(endpoint)
	configureAuthFunc(&reservationRecommendationDetailsClient.Client)

	reservationRecommendationsClient := reservationrecommendations.NewReservationRecommendationsClientWithBaseURI(endpoint)
	configureAuthFunc(&reservationRecommendationsClient.Client)

	reservationSummariesClient := reservationsummaries.NewReservationSummariesClientWithBaseURI(endpoint)
	configureAuthFunc(&reservationSummariesClient.Client)

	reservationTransactionsClient := reservationtransactions.NewReservationTransactionsClientWithBaseURI(endpoint)
	configureAuthFunc(&reservationTransactionsClient.Client)

	usageDetailsClient := usagedetails.NewUsageDetailsClientWithBaseURI(endpoint)
	configureAuthFunc(&usageDetailsClient.Client)

	return Client{
		AggregatedCost:                   &aggregatedCostClient,
		Balances:                         &balancesClient,
		Budgets:                          &budgetsClient,
		Charges:                          &chargesClient,
		Credits:                          &creditsClient,
		Events:                           &eventsClient,
		Lots:                             &lotsClient,
		Marketplaces:                     &marketplacesClient,
		PriceSheet:                       &priceSheetClient,
		ReservationDetails:               &reservationDetailsClient,
		ReservationRecommendationDetails: &reservationRecommendationDetailsClient,
		ReservationRecommendations:       &reservationRecommendationsClient,
		ReservationSummaries:             &reservationSummariesClient,
		ReservationTransactions:          &reservationTransactionsClient,
		UsageDetails:                     &usageDetailsClient,
	}
}
