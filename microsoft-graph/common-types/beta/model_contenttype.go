package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentType struct {
	AssociatedHubsUrls *[]string           `json:"associatedHubsUrls,omitempty"`
	Base               *ContentType        `json:"base,omitempty"`
	BaseTypes          *[]ContentType      `json:"baseTypes,omitempty"`
	ColumnLinks        *[]ColumnLink       `json:"columnLinks,omitempty"`
	ColumnPositions    *[]ColumnDefinition `json:"columnPositions,omitempty"`
	Columns            *[]ColumnDefinition `json:"columns,omitempty"`
	Description        *string             `json:"description,omitempty"`
	DocumentSet        *DocumentSet        `json:"documentSet,omitempty"`
	DocumentTemplate   *DocumentSetContent `json:"documentTemplate,omitempty"`
	Group              *string             `json:"group,omitempty"`
	Hidden             *bool               `json:"hidden,omitempty"`
	Id                 *string             `json:"id,omitempty"`
	InheritedFrom      *ItemReference      `json:"inheritedFrom,omitempty"`
	IsBuiltIn          *bool               `json:"isBuiltIn,omitempty"`
	Name               *string             `json:"name,omitempty"`
	ODataType          *string             `json:"@odata.type,omitempty"`
	Order              *ContentTypeOrder   `json:"order,omitempty"`
	ParentId           *string             `json:"parentId,omitempty"`
	PropagateChanges   *bool               `json:"propagateChanges,omitempty"`
	ReadOnly           *bool               `json:"readOnly,omitempty"`
	Sealed             *bool               `json:"sealed,omitempty"`
}
