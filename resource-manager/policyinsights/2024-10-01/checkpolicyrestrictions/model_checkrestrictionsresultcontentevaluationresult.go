package checkpolicyrestrictions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckRestrictionsResultContentEvaluationResult struct {
	PolicyEvaluations *[]PolicyEvaluationResult `json:"policyEvaluations,omitempty"`
}
