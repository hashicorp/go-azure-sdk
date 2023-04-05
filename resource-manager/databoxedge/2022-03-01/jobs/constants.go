package jobs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DownloadPhase string

const (
	DownloadPhaseDownloading  DownloadPhase = "Downloading"
	DownloadPhaseInitializing DownloadPhase = "Initializing"
	DownloadPhaseUnknown      DownloadPhase = "Unknown"
	DownloadPhaseVerifying    DownloadPhase = "Verifying"
)

func PossibleValuesForDownloadPhase() []string {
	return []string{
		string(DownloadPhaseDownloading),
		string(DownloadPhaseInitializing),
		string(DownloadPhaseUnknown),
		string(DownloadPhaseVerifying),
	}
}

type JobStatus string

const (
	JobStatusCanceled  JobStatus = "Canceled"
	JobStatusFailed    JobStatus = "Failed"
	JobStatusInvalid   JobStatus = "Invalid"
	JobStatusPaused    JobStatus = "Paused"
	JobStatusRunning   JobStatus = "Running"
	JobStatusScheduled JobStatus = "Scheduled"
	JobStatusSucceeded JobStatus = "Succeeded"
)

func PossibleValuesForJobStatus() []string {
	return []string{
		string(JobStatusCanceled),
		string(JobStatusFailed),
		string(JobStatusInvalid),
		string(JobStatusPaused),
		string(JobStatusRunning),
		string(JobStatusScheduled),
		string(JobStatusSucceeded),
	}
}

type JobType string

const (
	JobTypeBackup                JobType = "Backup"
	JobTypeDownloadUpdates       JobType = "DownloadUpdates"
	JobTypeInstallUpdates        JobType = "InstallUpdates"
	JobTypeInvalid               JobType = "Invalid"
	JobTypeRefreshContainer      JobType = "RefreshContainer"
	JobTypeRefreshShare          JobType = "RefreshShare"
	JobTypeRestore               JobType = "Restore"
	JobTypeScanForUpdates        JobType = "ScanForUpdates"
	JobTypeTriggerSupportPackage JobType = "TriggerSupportPackage"
)

func PossibleValuesForJobType() []string {
	return []string{
		string(JobTypeBackup),
		string(JobTypeDownloadUpdates),
		string(JobTypeInstallUpdates),
		string(JobTypeInvalid),
		string(JobTypeRefreshContainer),
		string(JobTypeRefreshShare),
		string(JobTypeRestore),
		string(JobTypeScanForUpdates),
		string(JobTypeTriggerSupportPackage),
	}
}

type UpdateOperationStage string

const (
	UpdateOperationStageDownloadComplete UpdateOperationStage = "DownloadComplete"
	UpdateOperationStageDownloadFailed   UpdateOperationStage = "DownloadFailed"
	UpdateOperationStageDownloadStarted  UpdateOperationStage = "DownloadStarted"
	UpdateOperationStageFailure          UpdateOperationStage = "Failure"
	UpdateOperationStageInitial          UpdateOperationStage = "Initial"
	UpdateOperationStageInstallComplete  UpdateOperationStage = "InstallComplete"
	UpdateOperationStageInstallFailed    UpdateOperationStage = "InstallFailed"
	UpdateOperationStageInstallStarted   UpdateOperationStage = "InstallStarted"
	UpdateOperationStageRebootInitiated  UpdateOperationStage = "RebootInitiated"
	UpdateOperationStageRescanComplete   UpdateOperationStage = "RescanComplete"
	UpdateOperationStageRescanFailed     UpdateOperationStage = "RescanFailed"
	UpdateOperationStageRescanStarted    UpdateOperationStage = "RescanStarted"
	UpdateOperationStageScanComplete     UpdateOperationStage = "ScanComplete"
	UpdateOperationStageScanFailed       UpdateOperationStage = "ScanFailed"
	UpdateOperationStageScanStarted      UpdateOperationStage = "ScanStarted"
	UpdateOperationStageSuccess          UpdateOperationStage = "Success"
	UpdateOperationStageUnknown          UpdateOperationStage = "Unknown"
)

func PossibleValuesForUpdateOperationStage() []string {
	return []string{
		string(UpdateOperationStageDownloadComplete),
		string(UpdateOperationStageDownloadFailed),
		string(UpdateOperationStageDownloadStarted),
		string(UpdateOperationStageFailure),
		string(UpdateOperationStageInitial),
		string(UpdateOperationStageInstallComplete),
		string(UpdateOperationStageInstallFailed),
		string(UpdateOperationStageInstallStarted),
		string(UpdateOperationStageRebootInitiated),
		string(UpdateOperationStageRescanComplete),
		string(UpdateOperationStageRescanFailed),
		string(UpdateOperationStageRescanStarted),
		string(UpdateOperationStageScanComplete),
		string(UpdateOperationStageScanFailed),
		string(UpdateOperationStageScanStarted),
		string(UpdateOperationStageSuccess),
		string(UpdateOperationStageUnknown),
	}
}
