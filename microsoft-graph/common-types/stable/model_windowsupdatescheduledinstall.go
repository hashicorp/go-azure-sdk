package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateScheduledInstall struct {
	ODataType            *string                                           `json:"@odata.type,omitempty"`
	ScheduledInstallDay  *WindowsUpdateScheduledInstallScheduledInstallDay `json:"scheduledInstallDay,omitempty"`
	ScheduledInstallTime *string                                           `json:"scheduledInstallTime,omitempty"`
}
