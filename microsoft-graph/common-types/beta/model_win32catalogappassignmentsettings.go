package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32CatalogAppAssignmentSettings struct {
	AutoUpdateSettings           *Win32LobAppAutoUpdateSettings                                 `json:"autoUpdateSettings,omitempty"`
	DeliveryOptimizationPriority *Win32CatalogAppAssignmentSettingsDeliveryOptimizationPriority `json:"deliveryOptimizationPriority,omitempty"`
	InstallTimeSettings          *MobileAppInstallTimeSettings                                  `json:"installTimeSettings,omitempty"`
	Notifications                *Win32CatalogAppAssignmentSettingsNotifications                `json:"notifications,omitempty"`
	ODataType                    *string                                                        `json:"@odata.type,omitempty"`
	RestartSettings              *Win32LobAppRestartSettings                                    `json:"restartSettings,omitempty"`
}
