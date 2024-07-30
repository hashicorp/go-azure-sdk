package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkDevice struct {
	Activity             *TeamworkDeviceActivity      `json:"activity,omitempty"`
	ActivityState        *TeamworkDeviceActivityState `json:"activityState,omitempty"`
	CompanyAssetTag      *string                      `json:"companyAssetTag,omitempty"`
	Configuration        *TeamworkDeviceConfiguration `json:"configuration,omitempty"`
	CreatedBy            *IdentitySet                 `json:"createdBy,omitempty"`
	CreatedDateTime      *string                      `json:"createdDateTime,omitempty"`
	CurrentUser          *TeamworkUserIdentity        `json:"currentUser,omitempty"`
	DeviceType           *TeamworkDeviceDeviceType    `json:"deviceType,omitempty"`
	HardwareDetail       *TeamworkHardwareDetail      `json:"hardwareDetail,omitempty"`
	Health               *TeamworkDeviceHealth        `json:"health,omitempty"`
	HealthStatus         *TeamworkDeviceHealthStatus  `json:"healthStatus,omitempty"`
	Id                   *string                      `json:"id,omitempty"`
	LastModifiedBy       *IdentitySet                 `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string                      `json:"lastModifiedDateTime,omitempty"`
	Notes                *string                      `json:"notes,omitempty"`
	ODataType            *string                      `json:"@odata.type,omitempty"`
	Operations           *[]TeamworkDeviceOperation   `json:"operations,omitempty"`
}
