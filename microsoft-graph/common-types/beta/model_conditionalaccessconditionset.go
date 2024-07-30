package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessConditionSet struct {
	Applications               *ConditionalAccessApplications                           `json:"applications,omitempty"`
	ClientAppTypes             *ConditionalAccessConditionSetClientAppTypes             `json:"clientAppTypes,omitempty"`
	ClientApplications         *ConditionalAccessClientApplications                     `json:"clientApplications,omitempty"`
	DeviceStates               *ConditionalAccessDeviceStates                           `json:"deviceStates,omitempty"`
	Devices                    *ConditionalAccessDevices                                `json:"devices,omitempty"`
	Locations                  *ConditionalAccessLocations                              `json:"locations,omitempty"`
	ODataType                  *string                                                  `json:"@odata.type,omitempty"`
	Platforms                  *ConditionalAccessPlatforms                              `json:"platforms,omitempty"`
	ServicePrincipalRiskLevels *ConditionalAccessConditionSetServicePrincipalRiskLevels `json:"servicePrincipalRiskLevels,omitempty"`
	SignInRiskLevels           *ConditionalAccessConditionSetSignInRiskLevels           `json:"signInRiskLevels,omitempty"`
	UserRiskLevels             *ConditionalAccessConditionSetUserRiskLevels             `json:"userRiskLevels,omitempty"`
	Users                      *ConditionalAccessUsers                                  `json:"users,omitempty"`
}
