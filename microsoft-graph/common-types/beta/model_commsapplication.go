package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommsApplication struct {
	Calls          *[]Call          `json:"calls,omitempty"`
	ODataType      *string          `json:"@odata.type,omitempty"`
	OnlineMeetings *[]OnlineMeeting `json:"onlineMeetings,omitempty"`
}
