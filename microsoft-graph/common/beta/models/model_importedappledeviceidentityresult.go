package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedAppleDeviceIdentityResult struct {
	CreatedDateTime                              *string                                           `json:"createdDateTime,omitempty"`
	Description                                  *string                                           `json:"description,omitempty"`
	DiscoverySource                              *ImportedAppleDeviceIdentityResultDiscoverySource `json:"discoverySource,omitempty"`
	EnrollmentState                              *ImportedAppleDeviceIdentityResultEnrollmentState `json:"enrollmentState,omitempty"`
	Id                                           *string                                           `json:"id,omitempty"`
	IsDeleted                                    *bool                                             `json:"isDeleted,omitempty"`
	IsSupervised                                 *bool                                             `json:"isSupervised,omitempty"`
	LastContactedDateTime                        *string                                           `json:"lastContactedDateTime,omitempty"`
	ODataType                                    *string                                           `json:"@odata.type,omitempty"`
	Platform                                     *ImportedAppleDeviceIdentityResultPlatform        `json:"platform,omitempty"`
	RequestedEnrollmentProfileAssignmentDateTime *string                                           `json:"requestedEnrollmentProfileAssignmentDateTime,omitempty"`
	RequestedEnrollmentProfileId                 *string                                           `json:"requestedEnrollmentProfileId,omitempty"`
	SerialNumber                                 *string                                           `json:"serialNumber,omitempty"`
	Status                                       *bool                                             `json:"status,omitempty"`
}
