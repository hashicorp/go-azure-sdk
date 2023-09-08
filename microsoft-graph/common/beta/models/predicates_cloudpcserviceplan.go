package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCServicePlanOperationPredicate struct {
	DisplayName     *string
	Id              *string
	ODataType       *string
	RamInGB         *int64
	StorageInGB     *int64
	UserProfileInGB *int64
	VCpuCount       *int64
}

func (p CloudPCServicePlanOperationPredicate) Matches(input CloudPCServicePlan) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RamInGB != nil && (input.RamInGB == nil || *p.RamInGB != *input.RamInGB) {
		return false
	}

	if p.StorageInGB != nil && (input.StorageInGB == nil || *p.StorageInGB != *input.StorageInGB) {
		return false
	}

	if p.UserProfileInGB != nil && (input.UserProfileInGB == nil || *p.UserProfileInGB != *input.UserProfileInGB) {
		return false
	}

	if p.VCpuCount != nil && (input.VCpuCount == nil || *p.VCpuCount != *input.VCpuCount) {
		return false
	}

	return true
}
