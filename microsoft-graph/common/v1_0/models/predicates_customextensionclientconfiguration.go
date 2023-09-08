package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomExtensionClientConfigurationOperationPredicate struct {
	ODataType             *string
	TimeoutInMilliseconds *int64
}

func (p CustomExtensionClientConfigurationOperationPredicate) Matches(input CustomExtensionClientConfiguration) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TimeoutInMilliseconds != nil && (input.TimeoutInMilliseconds == nil || *p.TimeoutInMilliseconds != *input.TimeoutInMilliseconds) {
		return false
	}

	return true
}
