package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintJob struct {
	Configuration   *PrintJobConfiguration `json:"configuration,omitempty"`
	CreatedBy       *UserIdentity          `json:"createdBy,omitempty"`
	CreatedDateTime *string                `json:"createdDateTime,omitempty"`
	Documents       *[]PrintDocument       `json:"documents,omitempty"`
	Id              *string                `json:"id,omitempty"`
	IsFetchable     *bool                  `json:"isFetchable,omitempty"`
	ODataType       *string                `json:"@odata.type,omitempty"`
	RedirectedFrom  *string                `json:"redirectedFrom,omitempty"`
	RedirectedTo    *string                `json:"redirectedTo,omitempty"`
	Status          *PrintJobStatus        `json:"status,omitempty"`
	Tasks           *[]PrintTask           `json:"tasks,omitempty"`
}
