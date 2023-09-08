package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppsAndServicesSettingsOperationPredicate struct {
	IsAppAndServicesTrialEnabled *bool
	IsOfficeStoreEnabled         *bool
	ODataType                    *string
}

func (p AppsAndServicesSettingsOperationPredicate) Matches(input AppsAndServicesSettings) bool {

	if p.IsAppAndServicesTrialEnabled != nil && (input.IsAppAndServicesTrialEnabled == nil || *p.IsAppAndServicesTrialEnabled != *input.IsAppAndServicesTrialEnabled) {
		return false
	}

	if p.IsOfficeStoreEnabled != nil && (input.IsOfficeStoreEnabled == nil || *p.IsOfficeStoreEnabled != *input.IsOfficeStoreEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
