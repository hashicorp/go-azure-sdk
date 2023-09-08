package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedEBook struct {
	Assignments           *[]ManagedEBookAssignment  `json:"assignments,omitempty"`
	CreatedDateTime       *string                    `json:"createdDateTime,omitempty"`
	Description           *string                    `json:"description,omitempty"`
	DeviceStates          *[]DeviceInstallState      `json:"deviceStates,omitempty"`
	DisplayName           *string                    `json:"displayName,omitempty"`
	Id                    *string                    `json:"id,omitempty"`
	InformationUrl        *string                    `json:"informationUrl,omitempty"`
	InstallSummary        *EBookInstallSummary       `json:"installSummary,omitempty"`
	LargeCover            *MimeContent               `json:"largeCover,omitempty"`
	LastModifiedDateTime  *string                    `json:"lastModifiedDateTime,omitempty"`
	ODataType             *string                    `json:"@odata.type,omitempty"`
	PrivacyInformationUrl *string                    `json:"privacyInformationUrl,omitempty"`
	PublishedDateTime     *string                    `json:"publishedDateTime,omitempty"`
	Publisher             *string                    `json:"publisher,omitempty"`
	UserStateSummary      *[]UserInstallStateSummary `json:"userStateSummary,omitempty"`
}
