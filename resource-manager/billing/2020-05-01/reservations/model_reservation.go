package reservations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Reservation struct {
	Id         *string                 `json:"id,omitempty"`
	Location   *string                 `json:"location,omitempty"`
	Name       *string                 `json:"name,omitempty"`
	Properties *ReservationProperty    `json:"properties,omitempty"`
	Sku        *ReservationSkuProperty `json:"sku,omitempty"`
	Type       *string                 `json:"type,omitempty"`
}
