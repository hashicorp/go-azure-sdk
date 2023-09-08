package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSynchronizationConnectionSettingsOperationPredicate struct {
	ClientId     *string
	ClientSecret *string
	ODataType    *string
}

func (p EducationSynchronizationConnectionSettingsOperationPredicate) Matches(input EducationSynchronizationConnectionSettings) bool {

	if p.ClientId != nil && (input.ClientId == nil || *p.ClientId != *input.ClientId) {
		return false
	}

	if p.ClientSecret != nil && (input.ClientSecret == nil || *p.ClientSecret != *input.ClientSecret) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
