package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageResourceAttribute struct {
	AttributeDestination           *AccessPackageResourceAttributeDestination `json:"attributeDestination,omitempty"`
	AttributeName                  *string                                    `json:"attributeName,omitempty"`
	AttributeSource                *AccessPackageResourceAttributeSource      `json:"attributeSource,omitempty"`
	Id                             *string                                    `json:"id,omitempty"`
	IsEditable                     *bool                                      `json:"isEditable,omitempty"`
	IsPersistedOnAssignmentRemoval *bool                                      `json:"isPersistedOnAssignmentRemoval,omitempty"`
	ODataType                      *string                                    `json:"@odata.type,omitempty"`
}
