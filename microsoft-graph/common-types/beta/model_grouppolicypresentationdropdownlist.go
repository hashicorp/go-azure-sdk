package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyPresentationDropdownList struct {
	DefaultItem          *GroupPolicyPresentationDropdownListItem   `json:"defaultItem,omitempty"`
	Definition           *GroupPolicyDefinition                     `json:"definition,omitempty"`
	Id                   *string                                    `json:"id,omitempty"`
	Items                *[]GroupPolicyPresentationDropdownListItem `json:"items,omitempty"`
	Label                *string                                    `json:"label,omitempty"`
	LastModifiedDateTime *string                                    `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                                    `json:"@odata.type,omitempty"`
	Required             *bool                                      `json:"required,omitempty"`
}
