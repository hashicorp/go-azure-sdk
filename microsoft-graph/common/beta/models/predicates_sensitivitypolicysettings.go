package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SensitivityPolicySettingsOperationPredicate struct {
	DowngradeSensitivityRequiresJustification *bool
	HelpWebUrl                                *string
	Id                                        *string
	IsMandatory                               *bool
	ODataType                                 *string
}

func (p SensitivityPolicySettingsOperationPredicate) Matches(input SensitivityPolicySettings) bool {

	if p.DowngradeSensitivityRequiresJustification != nil && (input.DowngradeSensitivityRequiresJustification == nil || *p.DowngradeSensitivityRequiresJustification != *input.DowngradeSensitivityRequiresJustification) {
		return false
	}

	if p.HelpWebUrl != nil && (input.HelpWebUrl == nil || *p.HelpWebUrl != *input.HelpWebUrl) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsMandatory != nil && (input.IsMandatory == nil || *p.IsMandatory != *input.IsMandatory) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
