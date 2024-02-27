package transfers

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TransferProperties struct {
	BillingAccountId       *string                   `json:"billingAccountId,omitempty"`
	BillingProfileId       *string                   `json:"billingProfileId,omitempty"`
	CanceledBy             *string                   `json:"canceledBy,omitempty"`
	CreationTime           *string                   `json:"creationTime,omitempty"`
	DetailedTransferStatus *[]DetailedTransferStatus `json:"detailedTransferStatus,omitempty"`
	ExpirationTime         *string                   `json:"expirationTime,omitempty"`
	InitiatorCustomerType  *string                   `json:"initiatorCustomerType,omitempty"`
	InitiatorEmailId       *string                   `json:"initiatorEmailId,omitempty"`
	InvoiceSectionId       *string                   `json:"invoiceSectionId,omitempty"`
	LastModifiedTime       *string                   `json:"lastModifiedTime,omitempty"`
	RecipientEmailId       *string                   `json:"recipientEmailId,omitempty"`
	ResellerId             *string                   `json:"resellerId,omitempty"`
	ResellerName           *string                   `json:"resellerName,omitempty"`
	TransferStatus         *TransferStatus           `json:"transferStatus,omitempty"`
}

func (o *TransferProperties) GetCreationTimeAsTime() (*time.Time, error) {
	if o.CreationTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreationTime, "2006-01-02T15:04:05Z07:00")
}

func (o *TransferProperties) GetExpirationTimeAsTime() (*time.Time, error) {
	if o.ExpirationTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ExpirationTime, "2006-01-02T15:04:05Z07:00")
}

func (o *TransferProperties) GetLastModifiedTimeAsTime() (*time.Time, error) {
	if o.LastModifiedTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastModifiedTime, "2006-01-02T15:04:05Z07:00")
}
