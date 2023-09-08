package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCHealthCheckItemOperationPredicate struct {
	AdditionalDetails       *string
	DisplayName             *string
	LastHealthCheckDateTime *string
	ODataType               *string
}

func (p CloudPCHealthCheckItemOperationPredicate) Matches(input CloudPCHealthCheckItem) bool {

	if p.AdditionalDetails != nil && (input.AdditionalDetails == nil || *p.AdditionalDetails != *input.AdditionalDetails) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.LastHealthCheckDateTime != nil && (input.LastHealthCheckDateTime == nil || *p.LastHealthCheckDateTime != *input.LastHealthCheckDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
