package billingrequest

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingRequestProperties struct {
	AdditionalInformation                *map[string]string    `json:"additionalInformation,omitempty"`
	BillingAccountDisplayName            *string               `json:"billingAccountDisplayName,omitempty"`
	BillingAccountId                     *string               `json:"billingAccountId,omitempty"`
	BillingAccountName                   *string               `json:"billingAccountName,omitempty"`
	BillingAccountPrimaryBillingTenantId *string               `json:"billingAccountPrimaryBillingTenantId,omitempty"`
	BillingProfileDisplayName            *string               `json:"billingProfileDisplayName,omitempty"`
	BillingProfileId                     *string               `json:"billingProfileId,omitempty"`
	BillingProfileName                   *string               `json:"billingProfileName,omitempty"`
	BillingScope                         *string               `json:"billingScope,omitempty"`
	CreatedBy                            *Principal            `json:"createdBy,omitempty"`
	CreationDate                         *string               `json:"creationDate,omitempty"`
	CustomerDisplayName                  *string               `json:"customerDisplayName,omitempty"`
	CustomerId                           *string               `json:"customerId,omitempty"`
	CustomerName                         *string               `json:"customerName,omitempty"`
	DecisionReason                       *string               `json:"decisionReason,omitempty"`
	ExpirationDate                       *string               `json:"expirationDate,omitempty"`
	InvoiceSectionDisplayName            *string               `json:"invoiceSectionDisplayName,omitempty"`
	InvoiceSectionId                     *string               `json:"invoiceSectionId,omitempty"`
	InvoiceSectionName                   *string               `json:"invoiceSectionName,omitempty"`
	Justification                        *string               `json:"justification,omitempty"`
	LastUpdatedBy                        *Principal            `json:"lastUpdatedBy,omitempty"`
	LastUpdatedDate                      *string               `json:"lastUpdatedDate,omitempty"`
	ProvisioningState                    *ProvisioningState    `json:"provisioningState,omitempty"`
	Recipients                           *[]Principal          `json:"recipients,omitempty"`
	RequestScope                         *string               `json:"requestScope,omitempty"`
	ReviewalDate                         *string               `json:"reviewalDate,omitempty"`
	ReviewedBy                           *Principal            `json:"reviewedBy,omitempty"`
	Status                               *BillingRequestStatus `json:"status,omitempty"`
	SubscriptionDisplayName              *string               `json:"subscriptionDisplayName,omitempty"`
	SubscriptionId                       *string               `json:"subscriptionId,omitempty"`
	SubscriptionName                     *string               `json:"subscriptionName,omitempty"`
	Type                                 *BillingRequestType   `json:"type,omitempty"`
}

func (o *BillingRequestProperties) GetCreationDateAsTime() (*time.Time, error) {
	if o.CreationDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreationDate, "2006-01-02T15:04:05Z07:00")
}

func (o *BillingRequestProperties) SetCreationDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreationDate = &formatted
}

func (o *BillingRequestProperties) GetExpirationDateAsTime() (*time.Time, error) {
	if o.ExpirationDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ExpirationDate, "2006-01-02T15:04:05Z07:00")
}

func (o *BillingRequestProperties) SetExpirationDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ExpirationDate = &formatted
}

func (o *BillingRequestProperties) GetLastUpdatedDateAsTime() (*time.Time, error) {
	if o.LastUpdatedDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastUpdatedDate, "2006-01-02T15:04:05Z07:00")
}

func (o *BillingRequestProperties) SetLastUpdatedDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastUpdatedDate = &formatted
}

func (o *BillingRequestProperties) GetReviewalDateAsTime() (*time.Time, error) {
	if o.ReviewalDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ReviewalDate, "2006-01-02T15:04:05Z07:00")
}

func (o *BillingRequestProperties) SetReviewalDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ReviewalDate = &formatted
}
