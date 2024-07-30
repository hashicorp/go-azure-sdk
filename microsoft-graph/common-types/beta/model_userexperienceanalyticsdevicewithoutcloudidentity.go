package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceWithoutCloudIdentity struct {
	AzureAdDeviceId *string `json:"azureAdDeviceId,omitempty"`
	DeviceName      *string `json:"deviceName,omitempty"`
	Id              *string `json:"id,omitempty"`
	ODataType       *string `json:"@odata.type,omitempty"`
}
