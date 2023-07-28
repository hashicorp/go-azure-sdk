package commitmenttiers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommitmentTierOperationPredicate struct {
	Kind     *string
	MaxCount *int64
	PlanType *string
	SkuName  *string
	Tier     *string
}

func (p CommitmentTierOperationPredicate) Matches(input CommitmentTier) bool {

	if p.Kind != nil && (input.Kind == nil || *p.Kind != *input.Kind) {
		return false
	}

	if p.MaxCount != nil && (input.MaxCount == nil || *p.MaxCount != *input.MaxCount) {
		return false
	}

	if p.PlanType != nil && (input.PlanType == nil || *p.PlanType != *input.PlanType) {
		return false
	}

	if p.SkuName != nil && (input.SkuName == nil || *p.SkuName != *input.SkuName) {
		return false
	}

	if p.Tier != nil && (input.Tier == nil || *p.Tier != *input.Tier) {
		return false
	}

	return true
}
