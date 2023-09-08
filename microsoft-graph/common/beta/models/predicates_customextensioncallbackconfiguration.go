package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomExtensionCallbackConfigurationOperationPredicate struct {
	ODataType       *string
	TimeoutDuration *string
}

func (p CustomExtensionCallbackConfigurationOperationPredicate) Matches(input CustomExtensionCallbackConfiguration) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TimeoutDuration != nil && (input.TimeoutDuration == nil || *p.TimeoutDuration != *input.TimeoutDuration) {
		return false
	}

	return true
}
