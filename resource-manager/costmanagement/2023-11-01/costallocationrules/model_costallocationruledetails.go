package costallocationrules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CostAllocationRuleDetails struct {
	SourceResources *[]SourceCostAllocationResource `json:"sourceResources,omitempty"`
	TargetResources *[]TargetCostAllocationResource `json:"targetResources,omitempty"`
}
