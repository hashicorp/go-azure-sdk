package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedDeviceIdentity struct {
	CreatedDateTime            *string                                           `json:"createdDateTime,omitempty"`
	Description                *string                                           `json:"description,omitempty"`
	EnrollmentState            *ImportedDeviceIdentityEnrollmentState            `json:"enrollmentState,omitempty"`
	Id                         *string                                           `json:"id,omitempty"`
	ImportedDeviceIdentifier   *string                                           `json:"importedDeviceIdentifier,omitempty"`
	ImportedDeviceIdentityType *ImportedDeviceIdentityImportedDeviceIdentityType `json:"importedDeviceIdentityType,omitempty"`
	LastContactedDateTime      *string                                           `json:"lastContactedDateTime,omitempty"`
	LastModifiedDateTime       *string                                           `json:"lastModifiedDateTime,omitempty"`
	ODataType                  *string                                           `json:"@odata.type,omitempty"`
	Platform                   *ImportedDeviceIdentityPlatform                   `json:"platform,omitempty"`
}
