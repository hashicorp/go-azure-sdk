package jobdefinitions

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CopyMode string

const (
	CopyModeAdditive CopyMode = "Additive"
	CopyModeMirror   CopyMode = "Mirror"
)

func PossibleValuesForCopyMode() []string {
	return []string{
		string(CopyModeAdditive),
		string(CopyModeMirror),
	}
}

func (s *CopyMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForCopyMode() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = CopyMode(decoded)
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
