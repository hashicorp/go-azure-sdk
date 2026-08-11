package netapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckAvailabilityResponse struct {
	IsAvailable *bool                     `json:"isAvailable,omitempty"`
	Message     *string                   `json:"message,omitempty"`
	Reason      *InAvailabilityReasonType `json:"reason,omitempty"`
}
