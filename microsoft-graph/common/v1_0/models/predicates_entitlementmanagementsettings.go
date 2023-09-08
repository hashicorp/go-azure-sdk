package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementSettingsOperationPredicate struct {
	DurationUntilExternalUserDeletedAfterBlocked *string
	Id                                           *string
	ODataType                                    *string
}

func (p EntitlementManagementSettingsOperationPredicate) Matches(input EntitlementManagementSettings) bool {

	if p.DurationUntilExternalUserDeletedAfterBlocked != nil && (input.DurationUntilExternalUserDeletedAfterBlocked == nil || *p.DurationUntilExternalUserDeletedAfterBlocked != *input.DurationUntilExternalUserDeletedAfterBlocked) {
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
