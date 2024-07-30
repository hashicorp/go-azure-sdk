package grouppolicyconfiguration

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyCategory struct {
	Children             *[]GroupPolicyCategory              `json:"children,omitempty"`
	DefinitionFile       *GroupPolicyDefinitionFile          `json:"definitionFile,omitempty"`
	Definitions          *[]GroupPolicyDefinition            `json:"definitions,omitempty"`
	DisplayName          *string                             `json:"displayName,omitempty"`
	Id                   *string                             `json:"id,omitempty"`
	IngestionSource      *GroupPolicyCategoryIngestionSource `json:"ingestionSource,omitempty"`
	IsRoot               *bool                               `json:"isRoot,omitempty"`
	LastModifiedDateTime *string                             `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                             `json:"@odata.type,omitempty"`
	Parent               *GroupPolicyCategory                `json:"parent,omitempty"`
}
