package usagesinformation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UsagesProperties struct {
	IsQuotaApplicable *bool         `json:"isQuotaApplicable,omitempty"`
	Name              *ResourceName `json:"name,omitempty"`
	Properties        *interface{}  `json:"properties,omitempty"`
	QuotaPeriod       *string       `json:"quotaPeriod,omitempty"`
	ResourceType      *string       `json:"resourceType,omitempty"`
	Unit              *string       `json:"unit,omitempty"`
	Usages            *UsagesObject `json:"usages,omitempty"`
}
