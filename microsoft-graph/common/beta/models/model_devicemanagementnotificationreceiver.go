package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementNotificationReceiver struct {
	ContactInformation *string `json:"contactInformation,omitempty"`
	Locale             *string `json:"locale,omitempty"`
	ODataType          *string `json:"@odata.type,omitempty"`
}
