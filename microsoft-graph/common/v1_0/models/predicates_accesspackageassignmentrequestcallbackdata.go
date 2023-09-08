package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentRequestCallbackDataOperationPredicate struct {
	CustomExtensionStageInstanceDetail *string
	CustomExtensionStageInstanceId     *string
	ODataType                          *string
	State                              *string
}

func (p AccessPackageAssignmentRequestCallbackDataOperationPredicate) Matches(input AccessPackageAssignmentRequestCallbackData) bool {

	if p.CustomExtensionStageInstanceDetail != nil && (input.CustomExtensionStageInstanceDetail == nil || *p.CustomExtensionStageInstanceDetail != *input.CustomExtensionStageInstanceDetail) {
		return false
	}

	if p.CustomExtensionStageInstanceId != nil && (input.CustomExtensionStageInstanceId == nil || *p.CustomExtensionStageInstanceId != *input.CustomExtensionStageInstanceId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.State != nil && (input.State == nil || *p.State != *input.State) {
		return false
	}

	return true
}
