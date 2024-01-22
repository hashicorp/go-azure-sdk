package v2021_10_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2021-10-01/billingsubscriptions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2021-10-01/billingsubscriptionsaliases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2021-10-01/paymentmethods"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	BillingSubscriptions        *billingsubscriptions.BillingSubscriptionsClient
	BillingSubscriptionsAliases *billingsubscriptionsaliases.BillingSubscriptionsAliasesClient
	PaymentMethods              *paymentmethods.PaymentMethodsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	billingSubscriptionsAliasesClient, err := billingsubscriptionsaliases.NewBillingSubscriptionsAliasesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BillingSubscriptionsAliases client: %+v", err)
	}
	configureFunc(billingSubscriptionsAliasesClient.Client)

	billingSubscriptionsClient, err := billingsubscriptions.NewBillingSubscriptionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BillingSubscriptions client: %+v", err)
	}
	configureFunc(billingSubscriptionsClient.Client)

	paymentMethodsClient, err := paymentmethods.NewPaymentMethodsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PaymentMethods client: %+v", err)
	}
	configureFunc(paymentMethodsClient.Client)

	return &Client{
		BillingSubscriptions:        billingSubscriptionsClient,
		BillingSubscriptionsAliases: billingSubscriptionsAliasesClient,
		PaymentMethods:              paymentMethodsClient,
	}, nil
}
