package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataUserMatchingSetting struct {
	MatchTarget      *IndustryDataUserMatchTargetReferenceValue `json:"matchTarget,omitempty"`
	ODataType        *string                                    `json:"@odata.type,omitempty"`
	PriorityOrder    *int64                                     `json:"priorityOrder,omitempty"`
	RoleGroup        *IndustryDataRoleGroup                     `json:"roleGroup,omitempty"`
	SourceIdentifier *IndustryDataIdentifierTypeReferenceValue  `json:"sourceIdentifier,omitempty"`
}
