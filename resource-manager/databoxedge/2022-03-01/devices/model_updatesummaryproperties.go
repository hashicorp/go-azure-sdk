package devices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateSummaryProperties struct {
	DeviceLastScannedDateTime            *string                `json:"deviceLastScannedDateTime,omitempty"`
	DeviceVersionNumber                  *string                `json:"deviceVersionNumber,omitempty"`
	FriendlyDeviceVersionName            *string                `json:"friendlyDeviceVersionName,omitempty"`
	InProgressDownloadJobId              *string                `json:"inProgressDownloadJobId,omitempty"`
	InProgressDownloadJobStartedDateTime *string                `json:"inProgressDownloadJobStartedDateTime,omitempty"`
	InProgressInstallJobId               *string                `json:"inProgressInstallJobId,omitempty"`
	InProgressInstallJobStartedDateTime  *string                `json:"inProgressInstallJobStartedDateTime,omitempty"`
	LastCompletedDownloadJobDateTime     *string                `json:"lastCompletedDownloadJobDateTime,omitempty"`
	LastCompletedDownloadJobId           *string                `json:"lastCompletedDownloadJobId,omitempty"`
	LastCompletedInstallJobDateTime      *string                `json:"lastCompletedInstallJobDateTime,omitempty"`
	LastCompletedInstallJobId            *string                `json:"lastCompletedInstallJobId,omitempty"`
	LastCompletedScanJobDateTime         *string                `json:"lastCompletedScanJobDateTime,omitempty"`
	LastDownloadJobStatus                *JobStatus             `json:"lastDownloadJobStatus,omitempty"`
	LastInstallJobStatus                 *JobStatus             `json:"lastInstallJobStatus,omitempty"`
	LastSuccessfulInstallJobDateTime     *string                `json:"lastSuccessfulInstallJobDateTime,omitempty"`
	LastSuccessfulScanJobTime            *string                `json:"lastSuccessfulScanJobTime,omitempty"`
	OngoingUpdateOperation               *UpdateOperation       `json:"ongoingUpdateOperation,omitempty"`
	RebootBehavior                       *InstallRebootBehavior `json:"rebootBehavior,omitempty"`
	TotalNumberOfUpdatesAvailable        *int64                 `json:"totalNumberOfUpdatesAvailable,omitempty"`
	TotalNumberOfUpdatesPendingDownload  *int64                 `json:"totalNumberOfUpdatesPendingDownload,omitempty"`
	TotalNumberOfUpdatesPendingInstall   *int64                 `json:"totalNumberOfUpdatesPendingInstall,omitempty"`
	TotalTimeInMinutes                   *int64                 `json:"totalTimeInMinutes,omitempty"`
	TotalUpdateSizeInBytes               *float64               `json:"totalUpdateSizeInBytes,omitempty"`
	UpdateTitles                         *[]string              `json:"updateTitles,omitempty"`
	Updates                              *[]UpdateDetails       `json:"updates,omitempty"`
}
