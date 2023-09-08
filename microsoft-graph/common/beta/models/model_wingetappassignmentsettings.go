package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WinGetAppAssignmentSettings struct {
	InstallTimeSettings *WinGetAppInstallTimeSettings             `json:"installTimeSettings,omitempty"`
	Notifications       *WinGetAppAssignmentSettingsNotifications `json:"notifications,omitempty"`
	ODataType           *string                                   `json:"@odata.type,omitempty"`
	RestartSettings     *WinGetAppRestartSettings                 `json:"restartSettings,omitempty"`
}
