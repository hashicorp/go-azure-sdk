package policyassignments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyAssignmentOperationPredicate struct {
	Name *string
}

func (p PolicyAssignmentOperationPredicate) Matches(input PolicyAssignment) bool {

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	return true
}
