package listtenantconfigurationviolations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ViolationOperationPredicate struct {
	ErrorMessage *string
	Id           *string
	UserId       *string
}

func (p ViolationOperationPredicate) Matches(input Violation) bool {

	if p.ErrorMessage != nil && (input.ErrorMessage == nil || *p.ErrorMessage != *input.ErrorMessage) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.UserId != nil && (input.UserId == nil || *p.UserId != *input.UserId) {
		return false
	}

	return true
}
