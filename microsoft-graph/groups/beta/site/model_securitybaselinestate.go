package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBaselineState struct {
	DisplayName                *string                         `json:"displayName,omitempty"`
	Id                         *string                         `json:"id,omitempty"`
	ODataType                  *string                         `json:"@odata.type,omitempty"`
	SecurityBaselineTemplateId *string                         `json:"securityBaselineTemplateId,omitempty"`
	SettingStates              *[]SecurityBaselineSettingState `json:"settingStates,omitempty"`
	State                      *SecurityBaselineStateState     `json:"state,omitempty"`
	UserPrincipalName          *string                         `json:"userPrincipalName,omitempty"`
}
