package netapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UsageResultOperationPredicate struct {
	Id *string
}

func (p UsageResultOperationPredicate) Matches(input UsageResult) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	return true
}
