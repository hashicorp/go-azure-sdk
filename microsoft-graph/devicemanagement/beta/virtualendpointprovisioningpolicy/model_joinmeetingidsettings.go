package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinMeetingIdSettings struct {
	IsPasscodeRequired *bool   `json:"isPasscodeRequired,omitempty"`
	JoinMeetingId      *string `json:"joinMeetingId,omitempty"`
	ODataType          *string `json:"@odata.type,omitempty"`
	Passcode           *string `json:"passcode,omitempty"`
}
