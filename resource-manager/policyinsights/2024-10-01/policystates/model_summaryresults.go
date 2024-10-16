package policystates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SummaryResults struct {
	NonCompliantPolicies  *int64              `json:"nonCompliantPolicies,omitempty"`
	NonCompliantResources *int64              `json:"nonCompliantResources,omitempty"`
	PolicyDetails         *[]ComplianceDetail `json:"policyDetails,omitempty"`
	PolicyGroupDetails    *[]ComplianceDetail `json:"policyGroupDetails,omitempty"`
	QueryResultsUri       *string             `json:"queryResultsUri,omitempty"`
	ResourceDetails       *[]ComplianceDetail `json:"resourceDetails,omitempty"`
}
