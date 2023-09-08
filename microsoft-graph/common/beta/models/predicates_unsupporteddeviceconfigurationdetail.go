package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnsupportedDeviceConfigurationDetailOperationPredicate struct {
	Message      *string
	ODataType    *string
	PropertyName *string
}

func (p UnsupportedDeviceConfigurationDetailOperationPredicate) Matches(input UnsupportedDeviceConfigurationDetail) bool {

	if p.Message != nil && (input.Message == nil || *p.Message != *input.Message) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PropertyName != nil && (input.PropertyName == nil || *p.PropertyName != *input.PropertyName) {
		return false
	}

	return true
}
