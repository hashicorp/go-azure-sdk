package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttributeMapping struct {
	DefaultValue            *string                       `json:"defaultValue,omitempty"`
	ExportMissingReferences *bool                         `json:"exportMissingReferences,omitempty"`
	FlowBehavior            *AttributeMappingFlowBehavior `json:"flowBehavior,omitempty"`
	FlowType                *AttributeMappingFlowType     `json:"flowType,omitempty"`
	MatchingPriority        *int64                        `json:"matchingPriority,omitempty"`
	ODataType               *string                       `json:"@odata.type,omitempty"`
	Source                  *AttributeMappingSource       `json:"source,omitempty"`
	TargetAttributeName     *string                       `json:"targetAttributeName,omitempty"`
}
