package capabilitiesbyserver

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CapabilityOperationPredicate struct {
	Name   *string
	Reason *string
}

func (p CapabilityOperationPredicate) Matches(input Capability) bool {

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Reason != nil && (input.Reason == nil || *p.Reason != *input.Reason) {
		return false
	}

	return true
}
