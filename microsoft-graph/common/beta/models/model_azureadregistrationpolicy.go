package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureADRegistrationPolicy struct {
	AllowedGroups       *[]string                           `json:"allowedGroups,omitempty"`
	AllowedUsers        *[]string                           `json:"allowedUsers,omitempty"`
	AppliesTo           *AzureADRegistrationPolicyAppliesTo `json:"appliesTo,omitempty"`
	IsAdminConfigurable *bool                               `json:"isAdminConfigurable,omitempty"`
	ODataType           *string                             `json:"@odata.type,omitempty"`
}
