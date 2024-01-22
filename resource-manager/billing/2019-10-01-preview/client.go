package v2019_10_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/address"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/agreements"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/availablebalances"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/billingaccounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/billingpermissions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/billingprofiles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/billingproperties"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/billingroleassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/billingroledefinitions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/billings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/billingsubscriptions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/customers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/departments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/enrollmentaccounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/instructions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/invoices"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/invoicesections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/paymentmethods"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/policies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/pricesheet"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/products"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/recipienttransfers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/transactions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/transfers"
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
	Departments            *departments.DepartmentsClient
	EnrollmentAccounts     *enrollmentaccounts.EnrollmentAccountsClient
	Instructions           *instructions.InstructionsClient
	InvoiceSections        *invoicesections.InvoiceSectionsClient
	Invoices               *invoices.InvoicesClient
	PaymentMethods         *paymentmethods.PaymentMethodsClient
	Policies               *policies.PoliciesClient
	PriceSheet             *pricesheet.PriceSheetClient
	Products               *products.ProductsClient
	RecipientTransfers     *recipienttransfers.RecipientTransfersClient
	Transactions           *transactions.TransactionsClient
	Transfers              *transfers.TransfersClient
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

	departmentsClient, err := departments.NewDepartmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Departments client: %+v", err)
	}
	configureFunc(departmentsClient.Client)

	enrollmentAccountsClient, err := enrollmentaccounts.NewEnrollmentAccountsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnrollmentAccounts client: %+v", err)
	}
	configureFunc(enrollmentAccountsClient.Client)

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

	paymentMethodsClient, err := paymentmethods.NewPaymentMethodsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PaymentMethods client: %+v", err)
	}
	configureFunc(paymentMethodsClient.Client)

	policiesClient, err := policies.NewPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Policies client: %+v", err)
	}
	configureFunc(policiesClient.Client)

	priceSheetClient, err := pricesheet.NewPriceSheetClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PriceSheet client: %+v", err)
	}
	configureFunc(priceSheetClient.Client)

	productsClient, err := products.NewProductsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Products client: %+v", err)
	}
	configureFunc(productsClient.Client)

	recipientTransfersClient, err := recipienttransfers.NewRecipientTransfersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RecipientTransfers client: %+v", err)
	}
	configureFunc(recipientTransfersClient.Client)

	transactionsClient, err := transactions.NewTransactionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Transactions client: %+v", err)
	}
	configureFunc(transactionsClient.Client)

	transfersClient, err := transfers.NewTransfersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Transfers client: %+v", err)
	}
	configureFunc(transfersClient.Client)

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
		Departments:            departmentsClient,
		EnrollmentAccounts:     enrollmentAccountsClient,
		Instructions:           instructionsClient,
		InvoiceSections:        invoiceSectionsClient,
		Invoices:               invoicesClient,
		PaymentMethods:         paymentMethodsClient,
		Policies:               policiesClient,
		PriceSheet:             priceSheetClient,
		Products:               productsClient,
		RecipientTransfers:     recipientTransfersClient,
		Transactions:           transactionsClient,
		Transfers:              transfersClient,
	}, nil
}
