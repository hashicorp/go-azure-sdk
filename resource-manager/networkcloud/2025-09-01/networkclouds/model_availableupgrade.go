package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AvailableUpgrade struct {
	AvailabilityLifecycle *AvailabilityLifecycle `json:"availabilityLifecycle,omitempty"`
	Version               *string                `json:"version,omitempty"`
}
