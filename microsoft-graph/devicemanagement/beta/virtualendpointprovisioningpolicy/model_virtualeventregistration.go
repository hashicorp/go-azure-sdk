package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventRegistration struct {
	CancelationDateTime         *string                                   `json:"cancelationDateTime,omitempty"`
	Email                       *string                                   `json:"email,omitempty"`
	FirstName                   *string                                   `json:"firstName,omitempty"`
	Id                          *string                                   `json:"id,omitempty"`
	LastName                    *string                                   `json:"lastName,omitempty"`
	ODataType                   *string                                   `json:"@odata.type,omitempty"`
	RegistrationDateTime        *string                                   `json:"registrationDateTime,omitempty"`
	RegistrationQuestionAnswers *[]VirtualEventRegistrationQuestionAnswer `json:"registrationQuestionAnswers,omitempty"`
	Sessions                    *[]VirtualEventSession                    `json:"sessions,omitempty"`
	Status                      *VirtualEventRegistrationStatus           `json:"status,omitempty"`
	UserId                      *string                                   `json:"userId,omitempty"`
}
