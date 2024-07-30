package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ZebraFotaDeploymentStatus struct {
	CancelRequested            *bool                               `json:"cancelRequested,omitempty"`
	CompleteOrCanceledDateTime *string                             `json:"completeOrCanceledDateTime,omitempty"`
	ErrorCode                  *ZebraFotaDeploymentStatusErrorCode `json:"errorCode,omitempty"`
	LastUpdatedDateTime        *string                             `json:"lastUpdatedDateTime,omitempty"`
	ODataType                  *string                             `json:"@odata.type,omitempty"`
	State                      *ZebraFotaDeploymentStatusState     `json:"state,omitempty"`
	TotalAwaitingInstall       *int64                              `json:"totalAwaitingInstall,omitempty"`
	TotalCanceled              *int64                              `json:"totalCanceled,omitempty"`
	TotalCreated               *int64                              `json:"totalCreated,omitempty"`
	TotalDevices               *int64                              `json:"totalDevices,omitempty"`
	TotalDownloading           *int64                              `json:"totalDownloading,omitempty"`
	TotalFailedDownload        *int64                              `json:"totalFailedDownload,omitempty"`
	TotalFailedInstall         *int64                              `json:"totalFailedInstall,omitempty"`
	TotalScheduled             *int64                              `json:"totalScheduled,omitempty"`
	TotalSucceededInstall      *int64                              `json:"totalSucceededInstall,omitempty"`
	TotalUnknown               *int64                              `json:"totalUnknown,omitempty"`
}
