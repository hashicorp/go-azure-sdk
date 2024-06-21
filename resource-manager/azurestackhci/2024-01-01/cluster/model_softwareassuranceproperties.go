package cluster

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SoftwareAssuranceProperties struct {
	LastUpdated             *string                  `json:"lastUpdated,omitempty"`
	SoftwareAssuranceIntent *SoftwareAssuranceIntent `json:"softwareAssuranceIntent,omitempty"`
	SoftwareAssuranceStatus *SoftwareAssuranceStatus `json:"softwareAssuranceStatus,omitempty"`
}
