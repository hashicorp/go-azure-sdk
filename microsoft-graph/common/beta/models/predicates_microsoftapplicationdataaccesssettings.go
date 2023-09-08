package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftApplicationDataAccessSettingsOperationPredicate struct {
	DisabledForGroup                     *string
	Id                                   *string
	IsEnabledForAllMicrosoftApplications *bool
	ODataType                            *string
}

func (p MicrosoftApplicationDataAccessSettingsOperationPredicate) Matches(input MicrosoftApplicationDataAccessSettings) bool {

	if p.DisabledForGroup != nil && (input.DisabledForGroup == nil || *p.DisabledForGroup != *input.DisabledForGroup) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsEnabledForAllMicrosoftApplications != nil && (input.IsEnabledForAllMicrosoftApplications == nil || *p.IsEnabledForAllMicrosoftApplications != *input.IsEnabledForAllMicrosoftApplications) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
