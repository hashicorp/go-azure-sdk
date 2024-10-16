package policystates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyGroupSummary struct {
	PolicyGroupName *string         `json:"policyGroupName,omitempty"`
	Results         *SummaryResults `json:"results,omitempty"`
}
