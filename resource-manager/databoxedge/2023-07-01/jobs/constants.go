package jobs

import (
	"encoding/json"
	"fmt"
	"strings"
)

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

func (s *DownloadPhase) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDownloadPhase(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDownloadPhase(input string) (*DownloadPhase, error) {
	vals := map[string]DownloadPhase{
		"downloading":  DownloadPhaseDownloading,
		"initializing": DownloadPhaseInitializing,
		"unknown":      DownloadPhaseUnknown,
		"verifying":    DownloadPhaseVerifying,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DownloadPhase(input)
	return &out, nil
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

func (s *JobStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseJobStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseJobStatus(input string) (*JobStatus, error) {
	vals := map[string]JobStatus{
		"canceled":  JobStatusCanceled,
		"failed":    JobStatusFailed,
		"invalid":   JobStatusInvalid,
		"paused":    JobStatusPaused,
		"running":   JobStatusRunning,
		"scheduled": JobStatusScheduled,
		"succeeded": JobStatusSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JobStatus(input)
	return &out, nil
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

func (s *JobType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseJobType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseJobType(input string) (*JobType, error) {
	vals := map[string]JobType{
		"backup":                JobTypeBackup,
		"downloadupdates":       JobTypeDownloadUpdates,
		"installupdates":        JobTypeInstallUpdates,
		"invalid":               JobTypeInvalid,
		"refreshcontainer":      JobTypeRefreshContainer,
		"refreshshare":          JobTypeRefreshShare,
		"restore":               JobTypeRestore,
		"scanforupdates":        JobTypeScanForUpdates,
		"triggersupportpackage": JobTypeTriggerSupportPackage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JobType(input)
	return &out, nil
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

func (s *UpdateOperationStage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUpdateOperationStage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUpdateOperationStage(input string) (*UpdateOperationStage, error) {
	vals := map[string]UpdateOperationStage{
		"downloadcomplete": UpdateOperationStageDownloadComplete,
		"downloadfailed":   UpdateOperationStageDownloadFailed,
		"downloadstarted":  UpdateOperationStageDownloadStarted,
		"failure":          UpdateOperationStageFailure,
		"initial":          UpdateOperationStageInitial,
		"installcomplete":  UpdateOperationStageInstallComplete,
		"installfailed":    UpdateOperationStageInstallFailed,
		"installstarted":   UpdateOperationStageInstallStarted,
		"rebootinitiated":  UpdateOperationStageRebootInitiated,
		"rescancomplete":   UpdateOperationStageRescanComplete,
		"rescanfailed":     UpdateOperationStageRescanFailed,
		"rescanstarted":    UpdateOperationStageRescanStarted,
		"scancomplete":     UpdateOperationStageScanComplete,
		"scanfailed":       UpdateOperationStageScanFailed,
		"scanstarted":      UpdateOperationStageScanStarted,
		"success":          UpdateOperationStageSuccess,
		"unknown":          UpdateOperationStageUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UpdateOperationStage(input)
	return &out, nil
}
