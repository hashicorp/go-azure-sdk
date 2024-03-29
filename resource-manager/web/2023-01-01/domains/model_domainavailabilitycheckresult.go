package domains

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainAvailabilityCheckResult struct {
	Available  *bool       `json:"available,omitempty"`
	DomainType *DomainType `json:"domainType,omitempty"`
	Name       *string     `json:"name,omitempty"`
}
