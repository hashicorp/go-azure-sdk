package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsRemoteConnection struct {
	CloudPCFailurePercentage *float64 `json:"cloudPcFailurePercentage,omitempty"`
	CloudPCRoundTripTime     *float64 `json:"cloudPcRoundTripTime,omitempty"`
	CloudPCSignInTime        *float64 `json:"cloudPcSignInTime,omitempty"`
	CoreBootTime             *float64 `json:"coreBootTime,omitempty"`
	CoreSignInTime           *float64 `json:"coreSignInTime,omitempty"`
	DeviceCount              *int64   `json:"deviceCount,omitempty"`
	DeviceId                 *string  `json:"deviceId,omitempty"`
	DeviceName               *string  `json:"deviceName,omitempty"`
	Id                       *string  `json:"id,omitempty"`
	Manufacturer             *string  `json:"manufacturer,omitempty"`
	Model                    *string  `json:"model,omitempty"`
	ODataType                *string  `json:"@odata.type,omitempty"`
	RemoteSignInTime         *float64 `json:"remoteSignInTime,omitempty"`
	UserPrincipalName        *string  `json:"userPrincipalName,omitempty"`
	VirtualNetwork           *string  `json:"virtualNetwork,omitempty"`
}
