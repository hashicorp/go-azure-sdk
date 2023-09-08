package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudCommunications struct {
	CallRecords    *[]CallRecordsCallRecord `json:"callRecords,omitempty"`
	Calls          *[]Call                  `json:"calls,omitempty"`
	ODataType      *string                  `json:"@odata.type,omitempty"`
	OnlineMeetings *[]OnlineMeeting         `json:"onlineMeetings,omitempty"`
	Presences      *[]Presence              `json:"presences,omitempty"`
}
