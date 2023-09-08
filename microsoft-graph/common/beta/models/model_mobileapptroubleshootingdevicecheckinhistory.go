package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppTroubleshootingDeviceCheckinHistory struct {
	ODataType                   *string                                      `json:"@odata.type,omitempty"`
	OccurrenceDateTime          *string                                      `json:"occurrenceDateTime,omitempty"`
	TroubleshootingErrorDetails *DeviceManagementTroubleshootingErrorDetails `json:"troubleshootingErrorDetails,omitempty"`
}
