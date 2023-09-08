package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceProvisioningLinkedResourceErrorDetailOperationPredicate struct {
	Code         *string
	Details      *string
	Message      *string
	ODataType    *string
	PropertyName *string
	Target       *string
}

func (p ServiceProvisioningLinkedResourceErrorDetailOperationPredicate) Matches(input ServiceProvisioningLinkedResourceErrorDetail) bool {

	if p.Code != nil && (input.Code == nil || *p.Code != *input.Code) {
		return false
	}

	if p.Details != nil && (input.Details == nil || *p.Details != *input.Details) {
		return false
	}

	if p.Message != nil && (input.Message == nil || *p.Message != *input.Message) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PropertyName != nil && (input.PropertyName == nil || *p.PropertyName != *input.PropertyName) {
		return false
	}

	if p.Target != nil && (input.Target == nil || *p.Target != *input.Target) {
		return false
	}

	return true
}
