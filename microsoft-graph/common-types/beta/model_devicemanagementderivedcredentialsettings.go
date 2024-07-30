package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementDerivedCredentialSettings struct {
	DisplayName                *string                                                    `json:"displayName,omitempty"`
	HelpUrl                    *string                                                    `json:"helpUrl,omitempty"`
	Id                         *string                                                    `json:"id,omitempty"`
	Issuer                     *DeviceManagementDerivedCredentialSettingsIssuer           `json:"issuer,omitempty"`
	NotificationType           *DeviceManagementDerivedCredentialSettingsNotificationType `json:"notificationType,omitempty"`
	ODataType                  *string                                                    `json:"@odata.type,omitempty"`
	RenewalThresholdPercentage *int64                                                     `json:"renewalThresholdPercentage,omitempty"`
}
