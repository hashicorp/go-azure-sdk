package marketplaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MarketplaceProperties struct {
	AccountName          *string  `json:"accountName,omitempty"`
	AdditionalInfo       *string  `json:"additionalInfo,omitempty"`
	AdditionalProperties *string  `json:"additionalProperties,omitempty"`
	BillingPeriodId      *string  `json:"billingPeriodId,omitempty"`
	ConsumedQuantity     *float64 `json:"consumedQuantity,omitempty"`
	ConsumedService      *string  `json:"consumedService,omitempty"`
	CostCenter           *string  `json:"costCenter,omitempty"`
	Currency             *string  `json:"currency,omitempty"`
	DepartmentName       *string  `json:"departmentName,omitempty"`
	InstanceId           *string  `json:"instanceId,omitempty"`
	InstanceName         *string  `json:"instanceName,omitempty"`
	IsEstimated          *bool    `json:"isEstimated,omitempty"`
	IsRecurringCharge    *bool    `json:"isRecurringCharge,omitempty"`
	MeterId              *string  `json:"meterId,omitempty"`
	OfferName            *string  `json:"offerName,omitempty"`
	OrderNumber          *string  `json:"orderNumber,omitempty"`
	PlanName             *string  `json:"planName,omitempty"`
	PretaxCost           *float64 `json:"pretaxCost,omitempty"`
	PublisherName        *string  `json:"publisherName,omitempty"`
	ResourceGroup        *string  `json:"resourceGroup,omitempty"`
	ResourceRate         *float64 `json:"resourceRate,omitempty"`
	SubscriptionGuid     *string  `json:"subscriptionGuid,omitempty"`
	SubscriptionName     *string  `json:"subscriptionName,omitempty"`
	UnitOfMeasure        *string  `json:"unitOfMeasure,omitempty"`
	UsageEnd             *string  `json:"usageEnd,omitempty"`
	UsageStart           *string  `json:"usageStart,omitempty"`
}
