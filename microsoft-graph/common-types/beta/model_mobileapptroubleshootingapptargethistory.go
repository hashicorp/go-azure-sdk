package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppTroubleshootingAppTargetHistory struct {
	ErrorCode                   *string                                           `json:"errorCode,omitempty"`
	ODataType                   *string                                           `json:"@odata.type,omitempty"`
	OccurrenceDateTime          *string                                           `json:"occurrenceDateTime,omitempty"`
	RunState                    *MobileAppTroubleshootingAppTargetHistoryRunState `json:"runState,omitempty"`
	SecurityGroupId             *string                                           `json:"securityGroupId,omitempty"`
	TroubleshootingErrorDetails *DeviceManagementTroubleshootingErrorDetails      `json:"troubleshootingErrorDetails,omitempty"`
}
