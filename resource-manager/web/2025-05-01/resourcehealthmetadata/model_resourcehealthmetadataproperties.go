package resourcehealthmetadata

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceHealthMetadataProperties struct {
	Category           *string `json:"category,omitempty"`
	SignalAvailability *bool   `json:"signalAvailability,omitempty"`
}
