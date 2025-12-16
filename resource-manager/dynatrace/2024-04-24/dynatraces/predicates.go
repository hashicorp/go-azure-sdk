package dynatraces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateResourceSupportedPropertiesOperationPredicate struct {
	CreationSupported *bool
	Name              *string
}

func (p CreateResourceSupportedPropertiesOperationPredicate) Matches(input CreateResourceSupportedProperties) bool {

	if p.CreationSupported != nil && (input.CreationSupported == nil || *p.CreationSupported != *input.CreationSupported) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	return true
}
