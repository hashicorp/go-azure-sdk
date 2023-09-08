package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomExtensionCalloutInstanceOperationPredicate struct {
	CustomExtensionId     *string
	Detail                *string
	ExternalCorrelationId *string
	Id                    *string
	ODataType             *string
}

func (p CustomExtensionCalloutInstanceOperationPredicate) Matches(input CustomExtensionCalloutInstance) bool {

	if p.CustomExtensionId != nil && (input.CustomExtensionId == nil || *p.CustomExtensionId != *input.CustomExtensionId) {
		return false
	}

	if p.Detail != nil && (input.Detail == nil || *p.Detail != *input.Detail) {
		return false
	}

	if p.ExternalCorrelationId != nil && (input.ExternalCorrelationId == nil || *p.ExternalCorrelationId != *input.ExternalCorrelationId) {
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
