package perimeterassociationproxies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkSecurityPerimeterConfigurationIssuesProperties struct {
	Description          *string                                             `json:"description,omitempty"`
	IssueType            *NetworkSecurityPerimeterConfigurationIssueType     `json:"issueType,omitempty"`
	Severity             *NetworkSecurityPerimeterConfigurationIssueSeverity `json:"severity,omitempty"`
	SuggestedAccessRules *[]string                                           `json:"suggestedAccessRules,omitempty"`
	SuggestedResourceIds *[]string                                           `json:"suggestedResourceIds,omitempty"`
}
