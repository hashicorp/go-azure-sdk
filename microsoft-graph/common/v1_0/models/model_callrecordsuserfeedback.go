package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsUserFeedback struct {
	ODataType *string                        `json:"@odata.type,omitempty"`
	Rating    *CallRecordsUserFeedbackRating `json:"rating,omitempty"`
	Text      *string                        `json:"text,omitempty"`
	Tokens    *CallRecordsFeedbackTokenSet   `json:"tokens,omitempty"`
}
