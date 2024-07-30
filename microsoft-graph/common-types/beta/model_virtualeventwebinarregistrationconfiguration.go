package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventWebinarRegistrationConfiguration struct {
	Capacity                *int64                                  `json:"capacity,omitempty"`
	Id                      *string                                 `json:"id,omitempty"`
	IsManualApprovalEnabled *bool                                   `json:"isManualApprovalEnabled,omitempty"`
	IsWaitlistEnabled       *bool                                   `json:"isWaitlistEnabled,omitempty"`
	ODataType               *string                                 `json:"@odata.type,omitempty"`
	Questions               *[]VirtualEventRegistrationQuestionBase `json:"questions,omitempty"`
	RegistrationWebUrl      *string                                 `json:"registrationWebUrl,omitempty"`
}
