package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AudioConferencing struct {
	ConferenceId    *string   `json:"conferenceId,omitempty"`
	DialinUrl       *string   `json:"dialinUrl,omitempty"`
	ODataType       *string   `json:"@odata.type,omitempty"`
	TollFreeNumber  *string   `json:"tollFreeNumber,omitempty"`
	TollFreeNumbers *[]string `json:"tollFreeNumbers,omitempty"`
	TollNumber      *string   `json:"tollNumber,omitempty"`
	TollNumbers     *[]string `json:"tollNumbers,omitempty"`
}
