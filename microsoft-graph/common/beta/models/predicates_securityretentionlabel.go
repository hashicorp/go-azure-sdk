package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRetentionLabelOperationPredicate struct {
	CreatedDateTime      *string
	DescriptionForAdmins *string
	DescriptionForUsers  *string
	DisplayName          *string
	Id                   *string
	IsInUse              *bool
	LabelToBeApplied     *string
	LastModifiedDateTime *string
	ODataType            *string
}

func (p SecurityRetentionLabelOperationPredicate) Matches(input SecurityRetentionLabel) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DescriptionForAdmins != nil && (input.DescriptionForAdmins == nil || *p.DescriptionForAdmins != *input.DescriptionForAdmins) {
		return false
	}

	if p.DescriptionForUsers != nil && (input.DescriptionForUsers == nil || *p.DescriptionForUsers != *input.DescriptionForUsers) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsInUse != nil && (input.IsInUse == nil || *p.IsInUse != *input.IsInUse) {
		return false
	}

	if p.LabelToBeApplied != nil && (input.LabelToBeApplied == nil || *p.LabelToBeApplied != *input.LabelToBeApplied) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
