package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InferenceDataOperationPredicate struct {
	ODataType               *string
	UserHasVerifiedAccuracy *bool
}

func (p InferenceDataOperationPredicate) Matches(input InferenceData) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UserHasVerifiedAccuracy != nil && (input.UserHasVerifiedAccuracy == nil || *p.UserHasVerifiedAccuracy != *input.UserHasVerifiedAccuracy) {
		return false
	}

	return true
}
