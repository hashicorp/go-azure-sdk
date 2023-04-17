package checkpolicyrestrictions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyEvaluationResult struct {
	EffectDetails     *PolicyEffectDetails               `json:"effectDetails,omitempty"`
	EvaluationDetails *CheckRestrictionEvaluationDetails `json:"evaluationDetails,omitempty"`
	EvaluationResult  *string                            `json:"evaluationResult,omitempty"`
	PolicyInfo        *PolicyReference                   `json:"policyInfo,omitempty"`
}
