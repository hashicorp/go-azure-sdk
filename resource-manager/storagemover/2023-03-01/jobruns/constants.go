package jobruns

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobRunScanStatus string

const (
	JobRunScanStatusCompleted  JobRunScanStatus = "Completed"
	JobRunScanStatusNotStarted JobRunScanStatus = "NotStarted"
	JobRunScanStatusScanning   JobRunScanStatus = "Scanning"
)

func PossibleValuesForJobRunScanStatus() []string {
	return []string{
		string(JobRunScanStatusCompleted),
		string(JobRunScanStatusNotStarted),
		string(JobRunScanStatusScanning),
	}
}

func (s *JobRunScanStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForJobRunScanStatus() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = JobRunScanStatus(decoded)
	return nil
}

type JobRunStatus string

const (
	JobRunStatusCancelRequested JobRunStatus = "CancelRequested"
	JobRunStatusCanceled        JobRunStatus = "Canceled"
	JobRunStatusCanceling       JobRunStatus = "Canceling"
	JobRunStatusFailed          JobRunStatus = "Failed"
	JobRunStatusQueued          JobRunStatus = "Queued"
	JobRunStatusRunning         JobRunStatus = "Running"
	JobRunStatusStarted         JobRunStatus = "Started"
	JobRunStatusSucceeded       JobRunStatus = "Succeeded"
)

func PossibleValuesForJobRunStatus() []string {
	return []string{
		string(JobRunStatusCancelRequested),
		string(JobRunStatusCanceled),
		string(JobRunStatusCanceling),
		string(JobRunStatusFailed),
		string(JobRunStatusQueued),
		string(JobRunStatusRunning),
		string(JobRunStatusStarted),
		string(JobRunStatusSucceeded),
	}
}

func (s *JobRunStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForJobRunStatus() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = JobRunStatus(decoded)
	return nil
}

type ProvisioningState string

const (
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateSucceeded),
	}
}

func (s *ProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForProvisioningState() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = ProvisioningState(decoded)
	return nil
}
