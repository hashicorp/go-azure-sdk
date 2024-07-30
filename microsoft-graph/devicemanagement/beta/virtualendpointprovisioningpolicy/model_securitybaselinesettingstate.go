package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBaselineSettingState struct {
	ContributingPolicies *[]SecurityBaselineContributingPolicy `json:"contributingPolicies,omitempty"`
	ErrorCode            *string                               `json:"errorCode,omitempty"`
	Id                   *string                               `json:"id,omitempty"`
	ODataType            *string                               `json:"@odata.type,omitempty"`
	SettingCategoryId    *string                               `json:"settingCategoryId,omitempty"`
	SettingCategoryName  *string                               `json:"settingCategoryName,omitempty"`
	SettingId            *string                               `json:"settingId,omitempty"`
	SettingName          *string                               `json:"settingName,omitempty"`
	SourcePolicies       *[]SettingSource                      `json:"sourcePolicies,omitempty"`
	State                *SecurityBaselineSettingStateState    `json:"state,omitempty"`
}
