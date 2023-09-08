package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsSegment struct {
	Callee        *CallRecordsEndpoint    `json:"callee,omitempty"`
	Caller        *CallRecordsEndpoint    `json:"caller,omitempty"`
	EndDateTime   *string                 `json:"endDateTime,omitempty"`
	FailureInfo   *CallRecordsFailureInfo `json:"failureInfo,omitempty"`
	Id            *string                 `json:"id,omitempty"`
	Media         *[]CallRecordsMedia     `json:"media,omitempty"`
	ODataType     *string                 `json:"@odata.type,omitempty"`
	StartDateTime *string                 `json:"startDateTime,omitempty"`
}
