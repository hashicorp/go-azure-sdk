package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewSet struct {
	Definitions        *[]AccessReviewScheduleDefinition `json:"definitions,omitempty"`
	HistoryDefinitions *[]AccessReviewHistoryDefinition  `json:"historyDefinitions,omitempty"`
	Id                 *string                           `json:"id,omitempty"`
	ODataType          *string                           `json:"@odata.type,omitempty"`
}
