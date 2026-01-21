package accounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImageInformationOperationPredicate struct {
	BatchSupportEndOfLife *string
	NodeAgentSKUId        *string
}

func (p ImageInformationOperationPredicate) Matches(input ImageInformation) bool {

	if p.BatchSupportEndOfLife != nil && (input.BatchSupportEndOfLife == nil || *p.BatchSupportEndOfLife != *input.BatchSupportEndOfLife) {
		return false
	}

	if p.NodeAgentSKUId != nil && *p.NodeAgentSKUId != input.NodeAgentSKUId {
		return false
	}

	return true
}

type PoolNodeCountsOperationPredicate struct {
	PoolId *string
}

func (p PoolNodeCountsOperationPredicate) Matches(input PoolNodeCounts) bool {

	if p.PoolId != nil && *p.PoolId != input.PoolId {
		return false
	}

	return true
}
