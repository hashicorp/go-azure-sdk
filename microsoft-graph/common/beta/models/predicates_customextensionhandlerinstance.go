package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomExtensionHandlerInstanceOperationPredicate struct {
	CustomExtensionId     *string
	ExternalCorrelationId *string
	ODataType             *string
}

func (p CustomExtensionHandlerInstanceOperationPredicate) Matches(input CustomExtensionHandlerInstance) bool {

	if p.CustomExtensionId != nil && (input.CustomExtensionId == nil || *p.CustomExtensionId != *input.CustomExtensionId) {
		return false
	}

	if p.ExternalCorrelationId != nil && (input.ExternalCorrelationId == nil || *p.ExternalCorrelationId != *input.ExternalCorrelationId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
