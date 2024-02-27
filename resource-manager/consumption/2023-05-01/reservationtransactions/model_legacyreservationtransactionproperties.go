package reservationtransactions

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LegacyReservationTransactionProperties struct {
	AccountName                *string  `json:"accountName,omitempty"`
	AccountOwnerEmail          *string  `json:"accountOwnerEmail,omitempty"`
	Amount                     *float64 `json:"amount,omitempty"`
	ArmSkuName                 *string  `json:"armSkuName,omitempty"`
	BillingFrequency           *string  `json:"billingFrequency,omitempty"`
	BillingMonth               *int64   `json:"billingMonth,omitempty"`
	CostCenter                 *string  `json:"costCenter,omitempty"`
	Currency                   *string  `json:"currency,omitempty"`
	CurrentEnrollment          *string  `json:"currentEnrollment,omitempty"`
	DepartmentName             *string  `json:"departmentName,omitempty"`
	Description                *string  `json:"description,omitempty"`
	EventDate                  *string  `json:"eventDate,omitempty"`
	EventType                  *string  `json:"eventType,omitempty"`
	MonetaryCommitment         *float64 `json:"monetaryCommitment,omitempty"`
	Overage                    *float64 `json:"overage,omitempty"`
	PurchasingEnrollment       *string  `json:"purchasingEnrollment,omitempty"`
	PurchasingSubscriptionGuid *string  `json:"purchasingSubscriptionGuid,omitempty"`
	PurchasingSubscriptionName *string  `json:"purchasingSubscriptionName,omitempty"`
	Quantity                   *float64 `json:"quantity,omitempty"`
	Region                     *string  `json:"region,omitempty"`
	ReservationOrderId         *string  `json:"reservationOrderId,omitempty"`
	ReservationOrderName       *string  `json:"reservationOrderName,omitempty"`
	Term                       *string  `json:"term,omitempty"`
}

func (o *LegacyReservationTransactionProperties) GetEventDateAsTime() (*time.Time, error) {
	if o.EventDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EventDate, "2006-01-02T15:04:05Z07:00")
}
