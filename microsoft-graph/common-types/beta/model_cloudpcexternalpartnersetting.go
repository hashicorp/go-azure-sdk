package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCExternalPartnerSetting struct {
	EnableConnection *bool                                `json:"enableConnection,omitempty"`
	Id               *string                              `json:"id,omitempty"`
	LastSyncDateTime *string                              `json:"lastSyncDateTime,omitempty"`
	ODataType        *string                              `json:"@odata.type,omitempty"`
	PartnerId        *string                              `json:"partnerId,omitempty"`
	Status           *CloudPCExternalPartnerSettingStatus `json:"status,omitempty"`
	StatusDetails    *string                              `json:"statusDetails,omitempty"`
}
