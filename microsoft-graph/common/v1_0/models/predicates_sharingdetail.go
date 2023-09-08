package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharingDetailOperationPredicate struct {
	ODataType      *string
	SharedDateTime *string
	SharingSubject *string
	SharingType    *string
}

func (p SharingDetailOperationPredicate) Matches(input SharingDetail) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SharedDateTime != nil && (input.SharedDateTime == nil || *p.SharedDateTime != *input.SharedDateTime) {
		return false
	}

	if p.SharingSubject != nil && (input.SharingSubject == nil || *p.SharingSubject != *input.SharingSubject) {
		return false
	}

	if p.SharingType != nil && (input.SharingType == nil || *p.SharingType != *input.SharingType) {
		return false
	}

	return true
}
