package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSynchronizationProfile struct {
	DataProvider                         *EducationSynchronizationDataProvider          `json:"dataProvider,omitempty"`
	DisplayName                          *string                                        `json:"displayName,omitempty"`
	Errors                               *[]EducationSynchronizationError               `json:"errors,omitempty"`
	ExpirationDate                       *string                                        `json:"expirationDate,omitempty"`
	HandleSpecialCharacterConstraint     *bool                                          `json:"handleSpecialCharacterConstraint,omitempty"`
	Id                                   *string                                        `json:"id,omitempty"`
	IdentitySynchronizationConfiguration *EducationIdentitySynchronizationConfiguration `json:"identitySynchronizationConfiguration,omitempty"`
	LicensesToAssign                     *[]EducationSynchronizationLicenseAssignment   `json:"licensesToAssign,omitempty"`
	ODataType                            *string                                        `json:"@odata.type,omitempty"`
	ProfileStatus                        *EducationSynchronizationProfileStatus         `json:"profileStatus,omitempty"`
	State                                *EducationSynchronizationProfileState          `json:"state,omitempty"`
}
