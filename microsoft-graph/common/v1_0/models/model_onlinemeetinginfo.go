package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingInfo struct {
	ConferenceId    *string   `json:"conferenceId,omitempty"`
	JoinUrl         *string   `json:"joinUrl,omitempty"`
	ODataType       *string   `json:"@odata.type,omitempty"`
	Phones          *[]Phone  `json:"phones,omitempty"`
	QuickDial       *string   `json:"quickDial,omitempty"`
	TollFreeNumbers *[]string `json:"tollFreeNumbers,omitempty"`
	TollNumber      *string   `json:"tollNumber,omitempty"`
}
