package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureResourceHealthSignal struct {
	Enabled    *ResourceHealthAvailabilityStateSignalBehavior `json:"enabled,omitempty"`
	SignalName *string                                        `json:"signalName,omitempty"`
	Status     *AzureResourceHealthSignalStatus               `json:"status,omitempty"`
}
