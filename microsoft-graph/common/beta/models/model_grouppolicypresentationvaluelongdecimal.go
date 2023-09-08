package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyPresentationValueLongDecimal struct {
	CreatedDateTime      *string                     `json:"createdDateTime,omitempty"`
	DefinitionValue      *GroupPolicyDefinitionValue `json:"definitionValue,omitempty"`
	Id                   *string                     `json:"id,omitempty"`
	LastModifiedDateTime *string                     `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                     `json:"@odata.type,omitempty"`
	Presentation         *GroupPolicyPresentation    `json:"presentation,omitempty"`
	Value                *int64                      `json:"value,omitempty"`
}
