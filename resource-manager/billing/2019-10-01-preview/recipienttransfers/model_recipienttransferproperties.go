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
	CreationTime           *string                   `json:"creationTime,omitempty"`
	DetailedTransferStatus *[]DetailedTransferStatus `json:"detailedTransferStatus,omitempty"`
	ExpirationTime         *string                   `json:"expirationTime,omitempty"`
	InitiatorCustomerType  *string                   `json:"initiatorCustomerType,omitempty"`
	InitiatorEmailId       *string                   `json:"initiatorEmailId,omitempty"`
	LastModifiedTime       *string                   `json:"lastModifiedTime,omitempty"`
	RecipientEmailId       *string                   `json:"recipientEmailId,omitempty"`
	ResellerId             *string                   `json:"resellerId,omitempty"`
	ResellerName           *string                   `json:"resellerName,omitempty"`
	TransferStatus         *TransferStatus           `json:"transferStatus,omitempty"`
}

func (o *RecipientTransferProperties) GetCreationTimeAsTime() (*time.Time, error) {
	if o.CreationTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreationTime, "2006-01-02T15:04:05Z07:00")
}

func (o *RecipientTransferProperties) SetCreationTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreationTime = &formatted
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

func (o *RecipientTransferProperties) GetLastModifiedTimeAsTime() (*time.Time, error) {
	if o.LastModifiedTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastModifiedTime, "2006-01-02T15:04:05Z07:00")
}

func (o *RecipientTransferProperties) SetLastModifiedTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastModifiedTime = &formatted
}
