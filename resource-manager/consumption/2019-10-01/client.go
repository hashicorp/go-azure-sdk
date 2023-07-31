package v2019_10_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2019-10-01/aggregatedcost"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2019-10-01/balances"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2019-10-01/budgets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2019-10-01/charges"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2019-10-01/credits"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2019-10-01/events"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2019-10-01/forecasts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2019-10-01/lots"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2019-10-01/marketplaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2019-10-01/pricesheet"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2019-10-01/reservationdetails"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2019-10-01/reservationrecommendationdetails"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2019-10-01/reservationrecommendations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2019-10-01/reservationsummaries"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2019-10-01/reservationtransactions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2019-10-01/usagedetails"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AggregatedCost                   *aggregatedcost.AggregatedCostClient
	Balances                         *balances.BalancesClient
	Budgets                          *budgets.BudgetsClient
	Charges                          *charges.ChargesClient
	Credits                          *credits.CreditsClient
	Events                           *events.EventsClient
	Forecasts                        *forecasts.ForecastsClient
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

func NewClientWithBaseURI(api sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	aggregatedCostClient, err := aggregatedcost.NewAggregatedCostClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building AggregatedCost client: %+v", err)
	}
	configureFunc(aggregatedCostClient.Client)

	balancesClient, err := balances.NewBalancesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Balances client: %+v", err)
	}
	configureFunc(balancesClient.Client)

	budgetsClient, err := budgets.NewBudgetsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Budgets client: %+v", err)
	}
	configureFunc(budgetsClient.Client)

	chargesClient, err := charges.NewChargesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Charges client: %+v", err)
	}
	configureFunc(chargesClient.Client)

	creditsClient, err := credits.NewCreditsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Credits client: %+v", err)
	}
	configureFunc(creditsClient.Client)

	eventsClient, err := events.NewEventsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Events client: %+v", err)
	}
	configureFunc(eventsClient.Client)

	forecastsClient, err := forecasts.NewForecastsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Forecasts client: %+v", err)
	}
	configureFunc(forecastsClient.Client)

	lotsClient, err := lots.NewLotsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Lots client: %+v", err)
	}
	configureFunc(lotsClient.Client)

	marketplacesClient, err := marketplaces.NewMarketplacesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Marketplaces client: %+v", err)
	}
	configureFunc(marketplacesClient.Client)

	priceSheetClient, err := pricesheet.NewPriceSheetClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PriceSheet client: %+v", err)
	}
	configureFunc(priceSheetClient.Client)

	reservationDetailsClient, err := reservationdetails.NewReservationDetailsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ReservationDetails client: %+v", err)
	}
	configureFunc(reservationDetailsClient.Client)

	reservationRecommendationDetailsClient, err := reservationrecommendationdetails.NewReservationRecommendationDetailsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ReservationRecommendationDetails client: %+v", err)
	}
	configureFunc(reservationRecommendationDetailsClient.Client)

	reservationRecommendationsClient, err := reservationrecommendations.NewReservationRecommendationsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ReservationRecommendations client: %+v", err)
	}
	configureFunc(reservationRecommendationsClient.Client)

	reservationSummariesClient, err := reservationsummaries.NewReservationSummariesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ReservationSummaries client: %+v", err)
	}
	configureFunc(reservationSummariesClient.Client)

	reservationTransactionsClient, err := reservationtransactions.NewReservationTransactionsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ReservationTransactions client: %+v", err)
	}
	configureFunc(reservationTransactionsClient.Client)

	usageDetailsClient, err := usagedetails.NewUsageDetailsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building UsageDetails client: %+v", err)
	}
	configureFunc(usageDetailsClient.Client)

	return &Client{
		AggregatedCost:                   aggregatedCostClient,
		Balances:                         balancesClient,
		Budgets:                          budgetsClient,
		Charges:                          chargesClient,
		Credits:                          creditsClient,
		Events:                           eventsClient,
		Forecasts:                        forecastsClient,
		Lots:                             lotsClient,
		Marketplaces:                     marketplacesClient,
		PriceSheet:                       priceSheetClient,
		ReservationDetails:               reservationDetailsClient,
		ReservationRecommendationDetails: reservationRecommendationDetailsClient,
		ReservationRecommendations:       reservationRecommendationsClient,
		ReservationSummaries:             reservationSummariesClient,
		ReservationTransactions:          reservationTransactionsClient,
		UsageDetails:                     usageDetailsClient,
	}, nil
}
