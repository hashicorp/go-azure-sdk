package policystates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Summary struct {
	OdataContext      *string                    `json:"@odata.context,omitempty"`
	OdataId           *string                    `json:"@odata.id,omitempty"`
	PolicyAssignments *[]PolicyAssignmentSummary `json:"policyAssignments,omitempty"`
	Results           *SummaryResults            `json:"results,omitempty"`
}
