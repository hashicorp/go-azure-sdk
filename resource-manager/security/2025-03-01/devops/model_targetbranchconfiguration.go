package devops

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetBranchConfiguration struct {
	AnnotateDefaultBranch *AnnotateDefaultBranchState `json:"annotateDefaultBranch,omitempty"`
	BranchNames           *[]string                   `json:"branchNames,omitempty"`
}
