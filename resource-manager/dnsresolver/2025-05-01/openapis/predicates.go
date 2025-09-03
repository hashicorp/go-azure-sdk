package openapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubResourceOperationPredicate struct {
	Id *string
}

func (p SubResourceOperationPredicate) Matches(input SubResource) bool {

	if p.Id != nil && *p.Id != input.Id {
		return false
	}

	return true
}

type VirtualNetworkDnsForwardingRulesetOperationPredicate struct {
	Id *string
}

func (p VirtualNetworkDnsForwardingRulesetOperationPredicate) Matches(input VirtualNetworkDnsForwardingRuleset) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	return true
}
