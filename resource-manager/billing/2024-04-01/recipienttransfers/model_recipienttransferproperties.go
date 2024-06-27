package recipienttransfers

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecipientTransferProperties struct {
	AllowedProductType     *[]EligibleProductType    `json:"allowedProductType,omitempty"`
	CanceledBy             *string                   `json:"canceledBy,omitempty"`
	CustomerTenantId       *string                   `json:"customerTenantId,omitempty"`
	DetailedTransferStatus *[]DetailedTransferStatus `json:"detailedTransferStatus,omitempty"`
	ExpirationTime         *string                   `json:"expirationTime,omitempty"`
	InitiatorCustomerType  *InitiatorCustomerType    `json:"initiatorCustomerType,omitempty"`
	InitiatorEmailId       *string                   `json:"initiatorEmailId,omitempty"`
	RecipientEmailId       *string                   `json:"recipientEmailId,omitempty"`
	ResellerId             *string                   `json:"resellerId,omitempty"`
	ResellerName           *string                   `json:"resellerName,omitempty"`
	SupportedAccounts      *[]SupportedAccountType   `json:"supportedAccounts,omitempty"`
	TransferStatus         *TransferStatus           `json:"transferStatus,omitempty"`
}

func (o *RecipientTransferProperties) GetExpirationTimeAsTime() (*time.Time, error) {
	if o.ExpirationTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ExpirationTime, "2006-01-02T15:04:05Z07:00")
}

func (o *RecipientTransferProperties) SetExpirationTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ExpirationTime = &formatted
}
