package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskAzureADGroupOperationPredicate struct {
	DisplayName *string
	GroupId     *string
	ODataType   *string
}

func (p WindowsKioskAzureADGroupOperationPredicate) Matches(input WindowsKioskAzureADGroup) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.GroupId != nil && (input.GroupId == nil || *p.GroupId != *input.GroupId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
