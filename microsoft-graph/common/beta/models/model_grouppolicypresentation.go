package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyPresentation struct {
	Definition           *GroupPolicyDefinition `json:"definition,omitempty"`
	Id                   *string                `json:"id,omitempty"`
	Label                *string                `json:"label,omitempty"`
	LastModifiedDateTime *string                `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                `json:"@odata.type,omitempty"`
}
