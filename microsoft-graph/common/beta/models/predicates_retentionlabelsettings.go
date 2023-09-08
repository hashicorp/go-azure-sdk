package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RetentionLabelSettingsOperationPredicate struct {
	IsContentUpdateAllowed  *bool
	IsDeleteAllowed         *bool
	IsLabelUpdateAllowed    *bool
	IsMetadataUpdateAllowed *bool
	IsRecordLocked          *bool
	ODataType               *string
}

func (p RetentionLabelSettingsOperationPredicate) Matches(input RetentionLabelSettings) bool {

	if p.IsContentUpdateAllowed != nil && (input.IsContentUpdateAllowed == nil || *p.IsContentUpdateAllowed != *input.IsContentUpdateAllowed) {
		return false
	}

	if p.IsDeleteAllowed != nil && (input.IsDeleteAllowed == nil || *p.IsDeleteAllowed != *input.IsDeleteAllowed) {
		return false
	}

	if p.IsLabelUpdateAllowed != nil && (input.IsLabelUpdateAllowed == nil || *p.IsLabelUpdateAllowed != *input.IsLabelUpdateAllowed) {
		return false
	}

	if p.IsMetadataUpdateAllowed != nil && (input.IsMetadataUpdateAllowed == nil || *p.IsMetadataUpdateAllowed != *input.IsMetadataUpdateAllowed) {
		return false
	}

	if p.IsRecordLocked != nil && (input.IsRecordLocked == nil || *p.IsRecordLocked != *input.IsRecordLocked) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
