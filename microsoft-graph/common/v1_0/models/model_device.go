package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Device struct {
	AccountEnabled                *bool                    `json:"accountEnabled,omitempty"`
	AlternativeSecurityIds        *[]AlternativeSecurityId `json:"alternativeSecurityIds,omitempty"`
	ApproximateLastSignInDateTime *string                  `json:"approximateLastSignInDateTime,omitempty"`
	ComplianceExpirationDateTime  *string                  `json:"complianceExpirationDateTime,omitempty"`
	DeletedDateTime               *string                  `json:"deletedDateTime,omitempty"`
	DeviceCategory                *string                  `json:"deviceCategory,omitempty"`
	DeviceId                      *string                  `json:"deviceId,omitempty"`
	DeviceMetadata                *string                  `json:"deviceMetadata,omitempty"`
	DeviceOwnership               *string                  `json:"deviceOwnership,omitempty"`
	DeviceVersion                 *int64                   `json:"deviceVersion,omitempty"`
	DisplayName                   *string                  `json:"displayName,omitempty"`
	EnrollmentProfileName         *string                  `json:"enrollmentProfileName,omitempty"`
	Extensions                    *[]Extension             `json:"extensions,omitempty"`
	Id                            *string                  `json:"id,omitempty"`
	IsCompliant                   *bool                    `json:"isCompliant,omitempty"`
	IsManaged                     *bool                    `json:"isManaged,omitempty"`
	MdmAppId                      *string                  `json:"mdmAppId,omitempty"`
	MemberOf                      *[]DirectoryObject       `json:"memberOf,omitempty"`
	ODataType                     *string                  `json:"@odata.type,omitempty"`
	OnPremisesLastSyncDateTime    *string                  `json:"onPremisesLastSyncDateTime,omitempty"`
	OnPremisesSyncEnabled         *bool                    `json:"onPremisesSyncEnabled,omitempty"`
	OperatingSystem               *string                  `json:"operatingSystem,omitempty"`
	OperatingSystemVersion        *string                  `json:"operatingSystemVersion,omitempty"`
	PhysicalIds                   *[]string                `json:"physicalIds,omitempty"`
	ProfileType                   *string                  `json:"profileType,omitempty"`
	RegisteredOwners              *[]DirectoryObject       `json:"registeredOwners,omitempty"`
	RegisteredUsers               *[]DirectoryObject       `json:"registeredUsers,omitempty"`
	RegistrationDateTime          *string                  `json:"registrationDateTime,omitempty"`
	SystemLabels                  *[]string                `json:"systemLabels,omitempty"`
	TransitiveMemberOf            *[]DirectoryObject       `json:"transitiveMemberOf,omitempty"`
	TrustType                     *string                  `json:"trustType,omitempty"`
}
