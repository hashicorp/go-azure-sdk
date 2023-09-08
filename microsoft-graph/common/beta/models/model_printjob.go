package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintJob struct {
	AcknowledgedDateTime *string                `json:"acknowledgedDateTime,omitempty"`
	CompletedDateTime    *string                `json:"completedDateTime,omitempty"`
	Configuration        *PrintJobConfiguration `json:"configuration,omitempty"`
	CreatedBy            *UserIdentity          `json:"createdBy,omitempty"`
	CreatedDateTime      *string                `json:"createdDateTime,omitempty"`
	DisplayName          *string                `json:"displayName,omitempty"`
	Documents            *[]PrintDocument       `json:"documents,omitempty"`
	ErrorCode            *int64                 `json:"errorCode,omitempty"`
	Id                   *string                `json:"id,omitempty"`
	IsFetchable          *bool                  `json:"isFetchable,omitempty"`
	ODataType            *string                `json:"@odata.type,omitempty"`
	RedirectedFrom       *string                `json:"redirectedFrom,omitempty"`
	RedirectedTo         *string                `json:"redirectedTo,omitempty"`
	Status               *PrintJobStatus        `json:"status,omitempty"`
	Tasks                *[]PrintTask           `json:"tasks,omitempty"`
}
