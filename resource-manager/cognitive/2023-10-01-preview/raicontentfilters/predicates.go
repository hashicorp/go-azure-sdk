package raicontentfilters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RaiContentFilterOperationPredicate struct {
	Description *string
	PolicyName  *string
}

func (p RaiContentFilterOperationPredicate) Matches(input RaiContentFilter) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.PolicyName != nil && (input.PolicyName == nil || *p.PolicyName != *input.PolicyName) {
		return false
	}

	return true
}
