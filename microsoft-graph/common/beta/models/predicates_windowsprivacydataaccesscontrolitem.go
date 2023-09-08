package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPrivacyDataAccessControlItemOperationPredicate struct {
	AppDisplayName       *string
	AppPackageFamilyName *string
	Id                   *string
	ODataType            *string
}

func (p WindowsPrivacyDataAccessControlItemOperationPredicate) Matches(input WindowsPrivacyDataAccessControlItem) bool {

	if p.AppDisplayName != nil && (input.AppDisplayName == nil || *p.AppDisplayName != *input.AppDisplayName) {
		return false
	}

	if p.AppPackageFamilyName != nil && (input.AppPackageFamilyName == nil || *p.AppPackageFamilyName != *input.AppPackageFamilyName) {
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
