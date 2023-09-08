package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommunicationsApplicationIdentityOperationPredicate struct {
	ApplicationType *string
	DisplayName     *string
	Hidden          *bool
	Id              *string
	ODataType       *string
}

func (p CommunicationsApplicationIdentityOperationPredicate) Matches(input CommunicationsApplicationIdentity) bool {

	if p.ApplicationType != nil && (input.ApplicationType == nil || *p.ApplicationType != *input.ApplicationType) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Hidden != nil && (input.Hidden == nil || *p.Hidden != *input.Hidden) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
