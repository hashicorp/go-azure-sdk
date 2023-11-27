package networksecurityperimeterconfigurations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NSPProvisioningIssueProperties struct {
	Description          *string   `json:"description,omitempty"`
	IssueType            *string   `json:"issueType,omitempty"`
	Severity             *string   `json:"severity,omitempty"`
	SuggestedAccessRules *[]string `json:"suggestedAccessRules,omitempty"`
	SuggestedResourceIds *[]string `json:"suggestedResourceIds,omitempty"`
}
