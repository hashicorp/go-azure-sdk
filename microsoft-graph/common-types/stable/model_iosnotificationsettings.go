package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosNotificationSettings struct {
	AlertType                *IosNotificationSettingsAlertType `json:"alertType,omitempty"`
	AppName                  *string                           `json:"appName,omitempty"`
	BadgesEnabled            *bool                             `json:"badgesEnabled,omitempty"`
	BundleID                 *string                           `json:"bundleID,omitempty"`
	Enabled                  *bool                             `json:"enabled,omitempty"`
	ODataType                *string                           `json:"@odata.type,omitempty"`
	Publisher                *string                           `json:"publisher,omitempty"`
	ShowInNotificationCenter *bool                             `json:"showInNotificationCenter,omitempty"`
	ShowOnLockScreen         *bool                             `json:"showOnLockScreen,omitempty"`
	SoundsEnabled            *bool                             `json:"soundsEnabled,omitempty"`
}
