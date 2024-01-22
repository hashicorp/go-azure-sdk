package v2020_05_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/address"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/agreements"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/availablebalances"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/billingaccounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/billingpermissions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/billingprofiles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/billingproperties"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/billingroleassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/billingroledefinitions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/billings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/billingsubscriptions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/customers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/instructions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/invoices"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/invoicesections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/policies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/products"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/reservations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/transactions"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Address                *address.AddressClient
	Agreements             *agreements.AgreementsClient
	AvailableBalances      *availablebalances.AvailableBalancesClient
	BillingAccounts        *billingaccounts.BillingAccountsClient
	BillingPermissions     *billingpermissions.BillingPermissionsClient
	BillingProfiles        *billingprofiles.BillingProfilesClient
	BillingProperties      *billingproperties.BillingPropertiesClient
	BillingRoleAssignments *billingroleassignments.BillingRoleAssignmentsClient
	BillingRoleDefinitions *billingroledefinitions.BillingRoleDefinitionsClient
	BillingSubscriptions   *billingsubscriptions.BillingSubscriptionsClient
	Billings               *billings.BillingsClient
	Customers              *customers.CustomersClient
	Instructions           *instructions.InstructionsClient
	InvoiceSections        *invoicesections.InvoiceSectionsClient
	Invoices               *invoices.InvoicesClient
	Policies               *policies.PoliciesClient
	Products               *products.ProductsClient
	Reservations           *reservations.ReservationsClient
	Transactions           *transactions.TransactionsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	addressClient, err := address.NewAddressClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Address client: %+v", err)
	}
	configureFunc(addressClient.Client)

	agreementsClient, err := agreements.NewAgreementsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Agreements client: %+v", err)
	}
	configureFunc(agreementsClient.Client)

	availableBalancesClient, err := availablebalances.NewAvailableBalancesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AvailableBalances client: %+v", err)
	}
	configureFunc(availableBalancesClient.Client)

	billingAccountsClient, err := billingaccounts.NewBillingAccountsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BillingAccounts client: %+v", err)
	}
	configureFunc(billingAccountsClient.Client)

	billingPermissionsClient, err := billingpermissions.NewBillingPermissionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BillingPermissions client: %+v", err)
	}
	configureFunc(billingPermissionsClient.Client)

	billingProfilesClient, err := billingprofiles.NewBillingProfilesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BillingProfiles client: %+v", err)
	}
	configureFunc(billingProfilesClient.Client)

	billingPropertiesClient, err := billingproperties.NewBillingPropertiesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BillingProperties client: %+v", err)
	}
	configureFunc(billingPropertiesClient.Client)

	billingRoleAssignmentsClient, err := billingroleassignments.NewBillingRoleAssignmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BillingRoleAssignments client: %+v", err)
	}
	configureFunc(billingRoleAssignmentsClient.Client)

	billingRoleDefinitionsClient, err := billingroledefinitions.NewBillingRoleDefinitionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BillingRoleDefinitions client: %+v", err)
	}
	configureFunc(billingRoleDefinitionsClient.Client)

	billingSubscriptionsClient, err := billingsubscriptions.NewBillingSubscriptionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BillingSubscriptions client: %+v", err)
	}
	configureFunc(billingSubscriptionsClient.Client)

	billingsClient, err := billings.NewBillingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Billings client: %+v", err)
	}
	configureFunc(billingsClient.Client)

	customersClient, err := customers.NewCustomersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Customers client: %+v", err)
	}
	configureFunc(customersClient.Client)

	instructionsClient, err := instructions.NewInstructionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Instructions client: %+v", err)
	}
	configureFunc(instructionsClient.Client)

	invoiceSectionsClient, err := invoicesections.NewInvoiceSectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InvoiceSections client: %+v", err)
	}
	configureFunc(invoiceSectionsClient.Client)

	invoicesClient, err := invoices.NewInvoicesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Invoices client: %+v", err)
	}
	configureFunc(invoicesClient.Client)

	policiesClient, err := policies.NewPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Policies client: %+v", err)
	}
	configureFunc(policiesClient.Client)

	productsClient, err := products.NewProductsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Products client: %+v", err)
	}
	configureFunc(productsClient.Client)

	reservationsClient, err := reservations.NewReservationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Reservations client: %+v", err)
	}
	configureFunc(reservationsClient.Client)

	transactionsClient, err := transactions.NewTransactionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Transactions client: %+v", err)
	}
	configureFunc(transactionsClient.Client)

	return &Client{
		Address:                addressClient,
		Agreements:             agreementsClient,
		AvailableBalances:      availableBalancesClient,
		BillingAccounts:        billingAccountsClient,
		BillingPermissions:     billingPermissionsClient,
		BillingProfiles:        billingProfilesClient,
		BillingProperties:      billingPropertiesClient,
		BillingRoleAssignments: billingRoleAssignmentsClient,
		BillingRoleDefinitions: billingRoleDefinitionsClient,
		BillingSubscriptions:   billingSubscriptionsClient,
		Billings:               billingsClient,
		Customers:              customersClient,
		Instructions:           instructionsClient,
		InvoiceSections:        invoiceSectionsClient,
		Invoices:               invoicesClient,
		Policies:               policiesClient,
		Products:               productsClient,
		Reservations:           reservationsClient,
		Transactions:           transactionsClient,
	}, nil
}
