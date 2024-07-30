package grouppolicyconfiguration

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type GroupPolicyConfigurationOperationPredicate struct {
}

func (p GroupPolicyConfigurationOperationPredicate) Matches(input beta.GroupPolicyConfiguration) bool {

	return true
}

type GroupPolicyConfigurationAssignmentOperationPredicate struct {
	Id                   *string
	LastModifiedDateTime *string
	ODataType            *string
}

func (p GroupPolicyConfigurationAssignmentOperationPredicate) Matches(input beta.GroupPolicyConfigurationAssignment) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
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
