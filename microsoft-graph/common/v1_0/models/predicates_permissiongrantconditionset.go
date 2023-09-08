package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionGrantConditionSetOperationPredicate struct {
	ClientApplicationsFromVerifiedPublisherOnly *bool
	Id                                          *string
	ODataType                                   *string
	PermissionClassification                    *string
	ResourceApplication                         *string
}

func (p PermissionGrantConditionSetOperationPredicate) Matches(input PermissionGrantConditionSet) bool {

	if p.ClientApplicationsFromVerifiedPublisherOnly != nil && (input.ClientApplicationsFromVerifiedPublisherOnly == nil || *p.ClientApplicationsFromVerifiedPublisherOnly != *input.ClientApplicationsFromVerifiedPublisherOnly) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PermissionClassification != nil && (input.PermissionClassification == nil || *p.PermissionClassification != *input.PermissionClassification) {
		return false
	}

	if p.ResourceApplication != nil && (input.ResourceApplication == nil || *p.ResourceApplication != *input.ResourceApplication) {
		return false
	}

	return true
}
