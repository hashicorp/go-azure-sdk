package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityUserFlowAttributeAssignment struct {
	DisplayName          *string                                           `json:"displayName,omitempty"`
	Id                   *string                                           `json:"id,omitempty"`
	IsOptional           *bool                                             `json:"isOptional,omitempty"`
	ODataType            *string                                           `json:"@odata.type,omitempty"`
	RequiresVerification *bool                                             `json:"requiresVerification,omitempty"`
	UserAttribute        *IdentityUserFlowAttribute                        `json:"userAttribute,omitempty"`
	UserAttributeValues  *[]UserAttributeValuesItem                        `json:"userAttributeValues,omitempty"`
	UserInputType        *IdentityUserFlowAttributeAssignmentUserInputType `json:"userInputType,omitempty"`
}
