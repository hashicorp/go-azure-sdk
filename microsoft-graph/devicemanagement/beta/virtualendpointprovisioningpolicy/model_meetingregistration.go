package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingRegistration struct {
	AllowedRegistrant         *MeetingRegistrationAllowedRegistrant `json:"allowedRegistrant,omitempty"`
	CustomQuestions           *[]MeetingRegistrationQuestion        `json:"customQuestions,omitempty"`
	Description               *string                               `json:"description,omitempty"`
	EndDateTime               *string                               `json:"endDateTime,omitempty"`
	Id                        *string                               `json:"id,omitempty"`
	ODataType                 *string                               `json:"@odata.type,omitempty"`
	Registrants               *[]MeetingRegistrantBase              `json:"registrants,omitempty"`
	RegistrationPageViewCount *int64                                `json:"registrationPageViewCount,omitempty"`
	RegistrationPageWebUrl    *string                               `json:"registrationPageWebUrl,omitempty"`
	Speakers                  *[]MeetingSpeaker                     `json:"speakers,omitempty"`
	StartDateTime             *string                               `json:"startDateTime,omitempty"`
	Subject                   *string                               `json:"subject,omitempty"`
}
