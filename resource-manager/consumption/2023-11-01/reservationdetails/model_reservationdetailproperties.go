package reservationdetails

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationDetailProperties struct {
	InstanceFlexibilityGroup *string  `json:"instanceFlexibilityGroup,omitempty"`
	InstanceFlexibilityRatio *string  `json:"instanceFlexibilityRatio,omitempty"`
	InstanceId               *string  `json:"instanceId,omitempty"`
	Kind                     *string  `json:"kind,omitempty"`
	ReservationId            *string  `json:"reservationId,omitempty"`
	ReservationOrderId       *string  `json:"reservationOrderId,omitempty"`
	ReservedHours            *float64 `json:"reservedHours,omitempty"`
	SkuName                  *string  `json:"skuName,omitempty"`
	TotalReservedQuantity    *float64 `json:"totalReservedQuantity,omitempty"`
	UsageDate                *string  `json:"usageDate,omitempty"`
	UsedHours                *float64 `json:"usedHours,omitempty"`
}
