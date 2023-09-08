package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementDerivedCredentialSettingsOperationPredicate struct {
	DisplayName                *string
	HelpUrl                    *string
	Id                         *string
	ODataType                  *string
	RenewalThresholdPercentage *int64
}

func (p DeviceManagementDerivedCredentialSettingsOperationPredicate) Matches(input DeviceManagementDerivedCredentialSettings) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.HelpUrl != nil && (input.HelpUrl == nil || *p.HelpUrl != *input.HelpUrl) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RenewalThresholdPercentage != nil && (input.RenewalThresholdPercentage == nil || *p.RenewalThresholdPercentage != *input.RenewalThresholdPercentage) {
		return false
	}

	return true
}
