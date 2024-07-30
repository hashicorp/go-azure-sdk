package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppTroubleshootingAppStateHistory struct {
	ActionType                  *MobileAppTroubleshootingAppStateHistoryActionType `json:"actionType,omitempty"`
	ErrorCode                   *string                                            `json:"errorCode,omitempty"`
	ODataType                   *string                                            `json:"@odata.type,omitempty"`
	OccurrenceDateTime          *string                                            `json:"occurrenceDateTime,omitempty"`
	RunState                    *MobileAppTroubleshootingAppStateHistoryRunState   `json:"runState,omitempty"`
	TroubleshootingErrorDetails *DeviceManagementTroubleshootingErrorDetails       `json:"troubleshootingErrorDetails,omitempty"`
}
