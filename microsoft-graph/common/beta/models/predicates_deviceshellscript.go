package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceShellScriptOperationPredicate struct {
	BlockExecutionNotifications *bool
	CreatedDateTime             *string
	Description                 *string
	DisplayName                 *string
	ExecutionFrequency          *string
	FileName                    *string
	Id                          *string
	LastModifiedDateTime        *string
	ODataType                   *string
	RetryCount                  *int64
	ScriptContent               *string
}

func (p DeviceShellScriptOperationPredicate) Matches(input DeviceShellScript) bool {

	if p.BlockExecutionNotifications != nil && (input.BlockExecutionNotifications == nil || *p.BlockExecutionNotifications != *input.BlockExecutionNotifications) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.ExecutionFrequency != nil && (input.ExecutionFrequency == nil || *p.ExecutionFrequency != *input.ExecutionFrequency) {
		return false
	}

	if p.FileName != nil && (input.FileName == nil || *p.FileName != *input.FileName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RetryCount != nil && (input.RetryCount == nil || *p.RetryCount != *input.RetryCount) {
		return false
	}

	if p.ScriptContent != nil && (input.ScriptContent == nil || *p.ScriptContent != *input.ScriptContent) {
		return false
	}

	return true
}
