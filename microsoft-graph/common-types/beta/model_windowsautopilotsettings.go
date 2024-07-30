package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotSettings struct {
	Id                            *string                             `json:"id,omitempty"`
	LastManualSyncTriggerDateTime *string                             `json:"lastManualSyncTriggerDateTime,omitempty"`
	LastSyncDateTime              *string                             `json:"lastSyncDateTime,omitempty"`
	ODataType                     *string                             `json:"@odata.type,omitempty"`
	SyncStatus                    *WindowsAutopilotSettingsSyncStatus `json:"syncStatus,omitempty"`
}
