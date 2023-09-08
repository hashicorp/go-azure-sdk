package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecordOperation struct {
	ClientContext        *string                          `json:"clientContext,omitempty"`
	CompletionReason     *RecordOperationCompletionReason `json:"completionReason,omitempty"`
	Id                   *string                          `json:"id,omitempty"`
	ODataType            *string                          `json:"@odata.type,omitempty"`
	RecordingAccessToken *string                          `json:"recordingAccessToken,omitempty"`
	RecordingLocation    *string                          `json:"recordingLocation,omitempty"`
	ResultInfo           *ResultInfo                      `json:"resultInfo,omitempty"`
	Status               *RecordOperationStatus           `json:"status,omitempty"`
}
