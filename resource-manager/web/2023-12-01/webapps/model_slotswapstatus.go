package webapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SlotSwapStatus struct {
	DestinationSlotName *string `json:"destinationSlotName,omitempty"`
	SourceSlotName      *string `json:"sourceSlotName,omitempty"`
	TimestampUtc        *string `json:"timestampUtc,omitempty"`
}
