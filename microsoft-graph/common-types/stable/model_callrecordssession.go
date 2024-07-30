package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsSession struct {
	Callee        *CallRecordsEndpoint          `json:"callee,omitempty"`
	Caller        *CallRecordsEndpoint          `json:"caller,omitempty"`
	EndDateTime   *string                       `json:"endDateTime,omitempty"`
	FailureInfo   *CallRecordsFailureInfo       `json:"failureInfo,omitempty"`
	Id            *string                       `json:"id,omitempty"`
	IsTest        *bool                         `json:"isTest,omitempty"`
	Modalities    *CallRecordsSessionModalities `json:"modalities,omitempty"`
	ODataType     *string                       `json:"@odata.type,omitempty"`
	Segments      *[]CallRecordsSegment         `json:"segments,omitempty"`
	StartDateTime *string                       `json:"startDateTime,omitempty"`
}
