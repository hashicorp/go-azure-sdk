package reservationorder

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationPurchaseRequest struct {
	Location   *string                               `json:"location,omitempty"`
	Properties *ReservationPurchaseRequestProperties `json:"properties,omitempty"`
	Sku        *SkuName                              `json:"sku,omitempty"`
}
