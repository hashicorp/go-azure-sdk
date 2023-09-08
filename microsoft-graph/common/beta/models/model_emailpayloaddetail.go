package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailPayloadDetail struct {
	Coachmarks       *[]PayloadCoachmark `json:"coachmarks,omitempty"`
	Content          *string             `json:"content,omitempty"`
	FromEmail        *string             `json:"fromEmail,omitempty"`
	FromName         *string             `json:"fromName,omitempty"`
	IsExternalSender *bool               `json:"isExternalSender,omitempty"`
	ODataType        *string             `json:"@odata.type,omitempty"`
	PhishingUrl      *string             `json:"phishingUrl,omitempty"`
	Subject          *string             `json:"subject,omitempty"`
}
