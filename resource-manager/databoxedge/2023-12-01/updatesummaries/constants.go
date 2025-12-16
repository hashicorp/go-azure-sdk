package updatesummaries

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InstallRebootBehavior string

const (
	InstallRebootBehaviorNeverReboots   InstallRebootBehavior = "NeverReboots"
	InstallRebootBehaviorRequestReboot  InstallRebootBehavior = "RequestReboot"
	InstallRebootBehaviorRequiresReboot InstallRebootBehavior = "RequiresReboot"
)

func PossibleValuesForInstallRebootBehavior() []string {
	return []string{
		string(InstallRebootBehaviorNeverReboots),
		string(InstallRebootBehaviorRequestReboot),
		string(InstallRebootBehaviorRequiresReboot),
	}
}

func (s *InstallRebootBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInstallRebootBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInstallRebootBehavior(input string) (*InstallRebootBehavior, error) {
	vals := map[string]InstallRebootBehavior{
		"neverreboots":   InstallRebootBehaviorNeverReboots,
		"requestreboot":  InstallRebootBehaviorRequestReboot,
		"requiresreboot": InstallRebootBehaviorRequiresReboot,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InstallRebootBehavior(input)
	return &out, nil
}

type InstallationImpact string

const (
	InstallationImpactDeviceRebooted          InstallationImpact = "DeviceRebooted"
	InstallationImpactKubernetesWorkloadsDown InstallationImpact = "KubernetesWorkloadsDown"
	InstallationImpactNone                    InstallationImpact = "None"
)

func PossibleValuesForInstallationImpact() []string {
	return []string{
		string(InstallationImpactDeviceRebooted),
		string(InstallationImpactKubernetesWorkloadsDown),
		string(InstallationImpactNone),
	}
}

func (s *InstallationImpact) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInstallationImpact(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInstallationImpact(input string) (*InstallationImpact, error) {
	vals := map[string]InstallationImpact{
		"devicerebooted":          InstallationImpactDeviceRebooted,
		"kubernetesworkloadsdown": InstallationImpactKubernetesWorkloadsDown,
		"none":                    InstallationImpactNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InstallationImpact(input)
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

type UpdateOperation string

const (
	UpdateOperationDownload UpdateOperation = "Download"
	UpdateOperationInstall  UpdateOperation = "Install"
	UpdateOperationNone     UpdateOperation = "None"
	UpdateOperationScan     UpdateOperation = "Scan"
)

func PossibleValuesForUpdateOperation() []string {
	return []string{
		string(UpdateOperationDownload),
		string(UpdateOperationInstall),
		string(UpdateOperationNone),
		string(UpdateOperationScan),
	}
}

func (s *UpdateOperation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUpdateOperation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUpdateOperation(input string) (*UpdateOperation, error) {
	vals := map[string]UpdateOperation{
		"download": UpdateOperationDownload,
		"install":  UpdateOperationInstall,
		"none":     UpdateOperationNone,
		"scan":     UpdateOperationScan,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UpdateOperation(input)
	return &out, nil
}

type UpdateStatus string

const (
	UpdateStatusDownloadCompleted UpdateStatus = "DownloadCompleted"
	UpdateStatusDownloadPending   UpdateStatus = "DownloadPending"
	UpdateStatusDownloadStarted   UpdateStatus = "DownloadStarted"
	UpdateStatusInstallCompleted  UpdateStatus = "InstallCompleted"
	UpdateStatusInstallStarted    UpdateStatus = "InstallStarted"
)

func PossibleValuesForUpdateStatus() []string {
	return []string{
		string(UpdateStatusDownloadCompleted),
		string(UpdateStatusDownloadPending),
		string(UpdateStatusDownloadStarted),
		string(UpdateStatusInstallCompleted),
		string(UpdateStatusInstallStarted),
	}
}

func (s *UpdateStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUpdateStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUpdateStatus(input string) (*UpdateStatus, error) {
	vals := map[string]UpdateStatus{
		"downloadcompleted": UpdateStatusDownloadCompleted,
		"downloadpending":   UpdateStatusDownloadPending,
		"downloadstarted":   UpdateStatusDownloadStarted,
		"installcompleted":  UpdateStatusInstallCompleted,
		"installstarted":    UpdateStatusInstallStarted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UpdateStatus(input)
	return &out, nil
}

type UpdateType string

const (
	UpdateTypeFirmware   UpdateType = "Firmware"
	UpdateTypeKubernetes UpdateType = "Kubernetes"
	UpdateTypeSoftware   UpdateType = "Software"
)

func PossibleValuesForUpdateType() []string {
	return []string{
		string(UpdateTypeFirmware),
		string(UpdateTypeKubernetes),
		string(UpdateTypeSoftware),
	}
}

func (s *UpdateType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUpdateType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUpdateType(input string) (*UpdateType, error) {
	vals := map[string]UpdateType{
		"firmware":   UpdateTypeFirmware,
		"kubernetes": UpdateTypeKubernetes,
		"software":   UpdateTypeSoftware,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UpdateType(input)
	return &out, nil
}
