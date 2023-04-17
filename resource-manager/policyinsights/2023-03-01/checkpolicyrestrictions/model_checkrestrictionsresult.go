package checkpolicyrestrictions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckRestrictionsResult struct {
	ContentEvaluationResult *CheckRestrictionsResultContentEvaluationResult `json:"contentEvaluationResult,omitempty"`
	FieldRestrictions       *[]FieldRestrictions                            `json:"fieldRestrictions,omitempty"`
}
