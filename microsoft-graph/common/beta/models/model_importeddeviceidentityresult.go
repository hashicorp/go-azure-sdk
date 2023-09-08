package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedDeviceIdentityResult struct {
	CreatedDateTime            *string                                                 `json:"createdDateTime,omitempty"`
	Description                *string                                                 `json:"description,omitempty"`
	EnrollmentState            *ImportedDeviceIdentityResultEnrollmentState            `json:"enrollmentState,omitempty"`
	Id                         *string                                                 `json:"id,omitempty"`
	ImportedDeviceIdentifier   *string                                                 `json:"importedDeviceIdentifier,omitempty"`
	ImportedDeviceIdentityType *ImportedDeviceIdentityResultImportedDeviceIdentityType `json:"importedDeviceIdentityType,omitempty"`
	LastContactedDateTime      *string                                                 `json:"lastContactedDateTime,omitempty"`
	LastModifiedDateTime       *string                                                 `json:"lastModifiedDateTime,omitempty"`
	ODataType                  *string                                                 `json:"@odata.type,omitempty"`
	Platform                   *ImportedDeviceIdentityResultPlatform                   `json:"platform,omitempty"`
	Status                     *bool                                                   `json:"status,omitempty"`
}
