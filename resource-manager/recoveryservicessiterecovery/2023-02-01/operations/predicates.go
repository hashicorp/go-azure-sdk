package operations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationsDiscoveryOperationPredicate struct {
	Name       *string
	Origin     *string
	Properties *interface{}
}

func (p OperationsDiscoveryOperationPredicate) Matches(input OperationsDiscovery) bool {

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	if p.Origin != nil && (input.Origin == nil && *p.Origin != *input.Origin) {
		return false
	}

	if p.Properties != nil && (input.Properties == nil && *p.Properties != *input.Properties) {
		return false
	}

	return true
}
