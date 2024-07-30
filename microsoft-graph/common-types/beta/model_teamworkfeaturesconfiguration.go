package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkFeaturesConfiguration struct {
	EmailToSendLogsAndFeedback   *string `json:"emailToSendLogsAndFeedback,omitempty"`
	IsAutoScreenShareEnabled     *bool   `json:"isAutoScreenShareEnabled,omitempty"`
	IsBluetoothBeaconingEnabled  *bool   `json:"isBluetoothBeaconingEnabled,omitempty"`
	IsHideMeetingNamesEnabled    *bool   `json:"isHideMeetingNamesEnabled,omitempty"`
	IsSendLogsAndFeedbackEnabled *bool   `json:"isSendLogsAndFeedbackEnabled,omitempty"`
	ODataType                    *string `json:"@odata.type,omitempty"`
}
