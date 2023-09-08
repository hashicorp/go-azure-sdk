package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingRegistrant struct {
	CustomQuestionAnswers *[]CustomQuestionAnswer  `json:"customQuestionAnswers,omitempty"`
	Email                 *string                  `json:"email,omitempty"`
	FirstName             *string                  `json:"firstName,omitempty"`
	Id                    *string                  `json:"id,omitempty"`
	JoinWebUrl            *string                  `json:"joinWebUrl,omitempty"`
	LastName              *string                  `json:"lastName,omitempty"`
	ODataType             *string                  `json:"@odata.type,omitempty"`
	RegistrationDateTime  *string                  `json:"registrationDateTime,omitempty"`
	Status                *MeetingRegistrantStatus `json:"status,omitempty"`
}
