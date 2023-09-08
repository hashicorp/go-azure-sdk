package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSettingsOperationPredicate struct {
	ContributionToContentDiscoveryAsOrganizationDisabled *bool
	ContributionToContentDiscoveryDisabled               *bool
	Id                                                   *string
	ODataType                                            *string
}

func (p UserSettingsOperationPredicate) Matches(input UserSettings) bool {

	if p.ContributionToContentDiscoveryAsOrganizationDisabled != nil && (input.ContributionToContentDiscoveryAsOrganizationDisabled == nil || *p.ContributionToContentDiscoveryAsOrganizationDisabled != *input.ContributionToContentDiscoveryAsOrganizationDisabled) {
		return false
	}

	if p.ContributionToContentDiscoveryDisabled != nil && (input.ContributionToContentDiscoveryDisabled == nil || *p.ContributionToContentDiscoveryDisabled != *input.ContributionToContentDiscoveryDisabled) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
