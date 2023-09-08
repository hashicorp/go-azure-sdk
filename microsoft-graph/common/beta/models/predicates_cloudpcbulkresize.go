package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCBulkResizeOperationPredicate struct {
	CreatedDateTime     *string
	DisplayName         *string
	Id                  *string
	ODataType           *string
	TargetServicePlanId *string
}

func (p CloudPCBulkResizeOperationPredicate) Matches(input CloudPCBulkResize) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TargetServicePlanId != nil && (input.TargetServicePlanId == nil || *p.TargetServicePlanId != *input.TargetServicePlanId) {
		return false
	}

	return true
}
