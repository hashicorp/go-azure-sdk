package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntityListResultOperationPredicate struct {
	Count    *int64
	NextLink *string
}

func (p EntityListResultOperationPredicate) Matches(input EntityListResult) bool {

	if p.Count != nil && (input.Count == nil || *p.Count != *input.Count) {
		return false
	}

	if p.NextLink != nil && (input.NextLink == nil || *p.NextLink != *input.NextLink) {
		return false
	}

	return true
}
