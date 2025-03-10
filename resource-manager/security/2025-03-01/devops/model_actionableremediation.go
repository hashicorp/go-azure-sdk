package devops

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActionableRemediation struct {
	BranchConfiguration    *TargetBranchConfiguration  `json:"branchConfiguration,omitempty"`
	CategoryConfigurations *[]CategoryConfiguration    `json:"categoryConfigurations,omitempty"`
	InheritFromParentState *InheritFromParentState     `json:"inheritFromParentState,omitempty"`
	State                  *ActionableRemediationState `json:"state,omitempty"`
}
