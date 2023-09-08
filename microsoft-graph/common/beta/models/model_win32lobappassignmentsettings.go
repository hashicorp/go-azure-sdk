package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppAssignmentSettings struct {
	DeliveryOptimizationPriority *Win32LobAppAssignmentSettingsDeliveryOptimizationPriority `json:"deliveryOptimizationPriority,omitempty"`
	InstallTimeSettings          *MobileAppInstallTimeSettings                              `json:"installTimeSettings,omitempty"`
	Notifications                *Win32LobAppAssignmentSettingsNotifications                `json:"notifications,omitempty"`
	ODataType                    *string                                                    `json:"@odata.type,omitempty"`
	RestartSettings              *Win32LobAppRestartSettings                                `json:"restartSettings,omitempty"`
}
