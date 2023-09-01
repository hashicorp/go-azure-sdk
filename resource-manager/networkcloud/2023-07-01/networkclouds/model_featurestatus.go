package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeatureStatus struct {
	DetailedStatus        *FeatureDetailedStatus `json:"detailedStatus,omitempty"`
	DetailedStatusMessage *string                `json:"detailedStatusMessage,omitempty"`
	Name                  *string                `json:"name,omitempty"`
	Version               *string                `json:"version,omitempty"`
}
