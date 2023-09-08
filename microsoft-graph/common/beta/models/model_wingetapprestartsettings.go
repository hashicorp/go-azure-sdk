package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WinGetAppRestartSettings struct {
	CountdownDisplayBeforeRestartInMinutes     *int64  `json:"countdownDisplayBeforeRestartInMinutes,omitempty"`
	GracePeriodInMinutes                       *int64  `json:"gracePeriodInMinutes,omitempty"`
	ODataType                                  *string `json:"@odata.type,omitempty"`
	RestartNotificationSnoozeDurationInMinutes *int64  `json:"restartNotificationSnoozeDurationInMinutes,omitempty"`
}
