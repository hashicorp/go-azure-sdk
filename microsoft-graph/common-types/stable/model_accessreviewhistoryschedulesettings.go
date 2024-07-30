package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewHistoryScheduleSettings struct {
	ODataType   *string              `json:"@odata.type,omitempty"`
	Recurrence  *PatternedRecurrence `json:"recurrence,omitempty"`
	ReportRange *string              `json:"reportRange,omitempty"`
}
