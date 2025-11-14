package v2024_08_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2024-08-01/budgets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2024-08-01/creditsummaries"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2024-08-01/openapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2024-08-01/pricesheet"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2024-08-01/pricesheetresults"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Budgets           *budgets.BudgetsClient
	CreditSummaries   *creditsummaries.CreditSummariesClient
	Openapis          *openapis.OpenapisClient
	PriceSheet        *pricesheet.PriceSheetClient
	PriceSheetResults *pricesheetresults.PriceSheetResultsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	budgetsClient, err := budgets.NewBudgetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Budgets client: %+v", err)
	}
	configureFunc(budgetsClient.Client)

	creditSummariesClient, err := creditsummaries.NewCreditSummariesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CreditSummaries client: %+v", err)
	}
	configureFunc(creditSummariesClient.Client)

	openapisClient, err := openapis.NewOpenapisClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Openapis client: %+v", err)
	}
	configureFunc(openapisClient.Client)

	priceSheetClient, err := pricesheet.NewPriceSheetClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PriceSheet client: %+v", err)
	}
	configureFunc(priceSheetClient.Client)

	priceSheetResultsClient, err := pricesheetresults.NewPriceSheetResultsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PriceSheetResults client: %+v", err)
	}
	configureFunc(priceSheetResultsClient.Client)

	return &Client{
		Budgets:           budgetsClient,
		CreditSummaries:   creditSummariesClient,
		Openapis:          openapisClient,
		PriceSheet:        priceSheetClient,
		PriceSheetResults: priceSheetResultsClient,
	}, nil
}
