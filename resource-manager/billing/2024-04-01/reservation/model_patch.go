package reservation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Patch struct {
	Properties *PatchProperties        `json:"properties,omitempty"`
	Sku        *ReservationSkuProperty `json:"sku,omitempty"`
	Tags       *map[string]string      `json:"tags,omitempty"`
}
