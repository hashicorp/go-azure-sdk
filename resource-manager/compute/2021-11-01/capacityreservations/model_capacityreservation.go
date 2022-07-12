package capacityreservations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CapacityReservation struct {
	Id         *string                        `json:"id,omitempty"`
	Location   string                         `json:"location"`
	Name       *string                        `json:"name,omitempty"`
	Properties *CapacityReservationProperties `json:"properties,omitempty"`
	Sku        Sku                            `json:"sku"`
	Tags       *map[string]string             `json:"tags,omitempty"`
	Type       *string                        `json:"type,omitempty"`
	Zones      *[]string                      `json:"zones,omitempty"`
}
