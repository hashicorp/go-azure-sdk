package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewSet struct {
	Decisions          *[]AccessReviewInstanceDecisionItem `json:"decisions,omitempty"`
	Definitions        *[]AccessReviewScheduleDefinition   `json:"definitions,omitempty"`
	HistoryDefinitions *[]AccessReviewHistoryDefinition    `json:"historyDefinitions,omitempty"`
	Id                 *string                             `json:"id,omitempty"`
	ODataType          *string                             `json:"@odata.type,omitempty"`
	Policy             *AccessReviewPolicy                 `json:"policy,omitempty"`
}
