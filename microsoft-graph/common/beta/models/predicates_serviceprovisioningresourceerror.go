package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceProvisioningResourceErrorOperationPredicate struct {
	CreatedDateTime *string
	IsResolved      *bool
	ODataType       *string
	ServiceInstance *string
}

func (p ServiceProvisioningResourceErrorOperationPredicate) Matches(input ServiceProvisioningResourceError) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.IsResolved != nil && (input.IsResolved == nil || *p.IsResolved != *input.IsResolved) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ServiceInstance != nil && (input.ServiceInstance == nil || *p.ServiceInstance != *input.ServiceInstance) {
		return false
	}

	return true
}
