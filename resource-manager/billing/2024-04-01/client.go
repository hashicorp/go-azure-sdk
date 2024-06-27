package v2024_04_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/agreement"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/associatedtenant"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/availablebalance"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/billingaccount"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/billingpermission"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/billingprofile"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/billingproperty"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/billingrequest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/billingroleassignment"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/billingroledefinition"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/billingsubscription"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/customer"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/department"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/enrollmentaccount"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/invoice"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/invoicesection"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/paymentmethods"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/policy"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/product"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/recipienttransfers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/reservation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/reservationorder"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/reservationorders"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/reservations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/savingsplan"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/savingsplanorder"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/transaction"
	"github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/transfers"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Agreement             *agreement.AgreementClient
	AssociatedTenant      *associatedtenant.AssociatedTenantClient
	AvailableBalance      *availablebalance.AvailableBalanceClient
	BillingAccount        *billingaccount.BillingAccountClient
	BillingPermission     *billingpermission.BillingPermissionClient
	BillingProfile        *billingprofile.BillingProfileClient
	BillingProperty       *billingproperty.BillingPropertyClient
	BillingRequest        *billingrequest.BillingRequestClient
	BillingRoleAssignment *billingroleassignment.BillingRoleAssignmentClient
	BillingRoleDefinition *billingroledefinition.BillingRoleDefinitionClient
	BillingSubscription   *billingsubscription.BillingSubscriptionClient
	Customer              *customer.CustomerClient
	Department            *department.DepartmentClient
	EnrollmentAccount     *enrollmentaccount.EnrollmentAccountClient
	Invoice               *invoice.InvoiceClient
	InvoiceSection        *invoicesection.InvoiceSectionClient
	PaymentMethods        *paymentmethods.PaymentMethodsClient
	Policy                *policy.PolicyClient
	Product               *product.ProductClient
	RecipientTransfers    *recipienttransfers.RecipientTransfersClient
	Reservation           *reservation.ReservationClient
	ReservationOrder      *reservationorder.ReservationOrderClient
	ReservationOrders     *reservationorders.ReservationOrdersClient
	Reservations          *reservations.ReservationsClient
	SavingsPlan           *savingsplan.SavingsPlanClient
	SavingsPlanOrder      *savingsplanorder.SavingsPlanOrderClient
	Transaction           *transaction.TransactionClient
	Transfers             *transfers.TransfersClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	agreementClient, err := agreement.NewAgreementClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Agreement client: %+v", err)
	}
	configureFunc(agreementClient.Client)

	associatedTenantClient, err := associatedtenant.NewAssociatedTenantClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AssociatedTenant client: %+v", err)
	}
	configureFunc(associatedTenantClient.Client)

	availableBalanceClient, err := availablebalance.NewAvailableBalanceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AvailableBalance client: %+v", err)
	}
	configureFunc(availableBalanceClient.Client)

	billingAccountClient, err := billingaccount.NewBillingAccountClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BillingAccount client: %+v", err)
	}
	configureFunc(billingAccountClient.Client)

	billingPermissionClient, err := billingpermission.NewBillingPermissionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BillingPermission client: %+v", err)
	}
	configureFunc(billingPermissionClient.Client)

	billingProfileClient, err := billingprofile.NewBillingProfileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BillingProfile client: %+v", err)
	}
	configureFunc(billingProfileClient.Client)

	billingPropertyClient, err := billingproperty.NewBillingPropertyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BillingProperty client: %+v", err)
	}
	configureFunc(billingPropertyClient.Client)

	billingRequestClient, err := billingrequest.NewBillingRequestClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BillingRequest client: %+v", err)
	}
	configureFunc(billingRequestClient.Client)

	billingRoleAssignmentClient, err := billingroleassignment.NewBillingRoleAssignmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BillingRoleAssignment client: %+v", err)
	}
	configureFunc(billingRoleAssignmentClient.Client)

	billingRoleDefinitionClient, err := billingroledefinition.NewBillingRoleDefinitionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BillingRoleDefinition client: %+v", err)
	}
	configureFunc(billingRoleDefinitionClient.Client)

	billingSubscriptionClient, err := billingsubscription.NewBillingSubscriptionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BillingSubscription client: %+v", err)
	}
	configureFunc(billingSubscriptionClient.Client)

	customerClient, err := customer.NewCustomerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Customer client: %+v", err)
	}
	configureFunc(customerClient.Client)

	departmentClient, err := department.NewDepartmentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Department client: %+v", err)
	}
	configureFunc(departmentClient.Client)

	enrollmentAccountClient, err := enrollmentaccount.NewEnrollmentAccountClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnrollmentAccount client: %+v", err)
	}
	configureFunc(enrollmentAccountClient.Client)

	invoiceClient, err := invoice.NewInvoiceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Invoice client: %+v", err)
	}
	configureFunc(invoiceClient.Client)

	invoiceSectionClient, err := invoicesection.NewInvoiceSectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InvoiceSection client: %+v", err)
	}
	configureFunc(invoiceSectionClient.Client)

	paymentMethodsClient, err := paymentmethods.NewPaymentMethodsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PaymentMethods client: %+v", err)
	}
	configureFunc(paymentMethodsClient.Client)

	policyClient, err := policy.NewPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Policy client: %+v", err)
	}
	configureFunc(policyClient.Client)

	productClient, err := product.NewProductClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Product client: %+v", err)
	}
	configureFunc(productClient.Client)

	recipientTransfersClient, err := recipienttransfers.NewRecipientTransfersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RecipientTransfers client: %+v", err)
	}
	configureFunc(recipientTransfersClient.Client)

	reservationClient, err := reservation.NewReservationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Reservation client: %+v", err)
	}
	configureFunc(reservationClient.Client)

	reservationOrderClient, err := reservationorder.NewReservationOrderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReservationOrder client: %+v", err)
	}
	configureFunc(reservationOrderClient.Client)

	reservationOrdersClient, err := reservationorders.NewReservationOrdersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReservationOrders client: %+v", err)
	}
	configureFunc(reservationOrdersClient.Client)

	reservationsClient, err := reservations.NewReservationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Reservations client: %+v", err)
	}
	configureFunc(reservationsClient.Client)

	savingsPlanClient, err := savingsplan.NewSavingsPlanClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SavingsPlan client: %+v", err)
	}
	configureFunc(savingsPlanClient.Client)

	savingsPlanOrderClient, err := savingsplanorder.NewSavingsPlanOrderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SavingsPlanOrder client: %+v", err)
	}
	configureFunc(savingsPlanOrderClient.Client)

	transactionClient, err := transaction.NewTransactionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Transaction client: %+v", err)
	}
	configureFunc(transactionClient.Client)

	transfersClient, err := transfers.NewTransfersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Transfers client: %+v", err)
	}
	configureFunc(transfersClient.Client)

	return &Client{
		Agreement:             agreementClient,
		AssociatedTenant:      associatedTenantClient,
		AvailableBalance:      availableBalanceClient,
		BillingAccount:        billingAccountClient,
		BillingPermission:     billingPermissionClient,
		BillingProfile:        billingProfileClient,
		BillingProperty:       billingPropertyClient,
		BillingRequest:        billingRequestClient,
		BillingRoleAssignment: billingRoleAssignmentClient,
		BillingRoleDefinition: billingRoleDefinitionClient,
		BillingSubscription:   billingSubscriptionClient,
		Customer:              customerClient,
		Department:            departmentClient,
		EnrollmentAccount:     enrollmentAccountClient,
		Invoice:               invoiceClient,
		InvoiceSection:        invoiceSectionClient,
		PaymentMethods:        paymentMethodsClient,
		Policy:                policyClient,
		Product:               productClient,
		RecipientTransfers:    recipientTransfersClient,
		Reservation:           reservationClient,
		ReservationOrder:      reservationOrderClient,
		ReservationOrders:     reservationOrdersClient,
		Reservations:          reservationsClient,
		SavingsPlan:           savingsPlanClient,
		SavingsPlanOrder:      savingsPlanOrderClient,
		Transaction:           transactionClient,
		Transfers:             transfersClient,
	}, nil
}
