package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationManagerClientHealthStateOperationPredicate struct {
	ErrorCode        *int64
	LastSyncDateTime *string
	ODataType        *string
}

func (p ConfigurationManagerClientHealthStateOperationPredicate) Matches(input ConfigurationManagerClientHealthState) bool {

	if p.ErrorCode != nil && (input.ErrorCode == nil || *p.ErrorCode != *input.ErrorCode) {
		return false
	}

	if p.LastSyncDateTime != nil && (input.LastSyncDateTime == nil || *p.LastSyncDateTime != *input.LastSyncDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
