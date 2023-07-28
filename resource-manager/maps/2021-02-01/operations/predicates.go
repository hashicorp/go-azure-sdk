package operations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationDetailOperationPredicate struct {
	IsDataAction *bool
	Name         *string
	Origin       *string
}

func (p OperationDetailOperationPredicate) Matches(input OperationDetail) bool {

	if p.IsDataAction != nil && (input.IsDataAction == nil || *p.IsDataAction != *input.IsDataAction) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Origin != nil && (input.Origin == nil || *p.Origin != *input.Origin) {
		return false
	}

	return true
}
