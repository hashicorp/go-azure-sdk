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
	out, err := parseCopyMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCopyMode(input string) (*CopyMode, error) {
	vals := map[string]CopyMode{
		"additive": CopyModeAdditive,
		"mirror":   CopyModeMirror,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CopyMode(input)
	return &out, nil
}

type DataIntegrityValidation string

const (
	DataIntegrityValidationNone                 DataIntegrityValidation = "None"
	DataIntegrityValidationSaveFileMDFive       DataIntegrityValidation = "SaveFileMD5"
	DataIntegrityValidationSaveVerifyFileMDFive DataIntegrityValidation = "SaveVerifyFileMD5"
)

func PossibleValuesForDataIntegrityValidation() []string {
	return []string{
		string(DataIntegrityValidationNone),
		string(DataIntegrityValidationSaveFileMDFive),
		string(DataIntegrityValidationSaveVerifyFileMDFive),
	}
}

func (s *DataIntegrityValidation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDataIntegrityValidation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDataIntegrityValidation(input string) (*DataIntegrityValidation, error) {
	vals := map[string]DataIntegrityValidation{
		"none":              DataIntegrityValidationNone,
		"savefilemd5":       DataIntegrityValidationSaveFileMDFive,
		"saveverifyfilemd5": DataIntegrityValidationSaveVerifyFileMDFive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataIntegrityValidation(input)
	return &out, nil
}

type Frequency string

const (
	FrequencyDaily   Frequency = "Daily"
	FrequencyHourly  Frequency = "Hourly"
	FrequencyMonthly Frequency = "Monthly"
	FrequencyNone    Frequency = "None"
	FrequencyOnetime Frequency = "Onetime"
	FrequencyWeekly  Frequency = "Weekly"
)

func PossibleValuesForFrequency() []string {
	return []string{
		string(FrequencyDaily),
		string(FrequencyHourly),
		string(FrequencyMonthly),
		string(FrequencyNone),
		string(FrequencyOnetime),
		string(FrequencyWeekly),
	}
}

func (s *Frequency) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFrequency(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFrequency(input string) (*Frequency, error) {
	vals := map[string]Frequency{
		"daily":   FrequencyDaily,
		"hourly":  FrequencyHourly,
		"monthly": FrequencyMonthly,
		"none":    FrequencyNone,
		"onetime": FrequencyOnetime,
		"weekly":  FrequencyWeekly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Frequency(input)
	return &out, nil
}

type JobRunStatus string

const (
	JobRunStatusCancelRequested             JobRunStatus = "CancelRequested"
	JobRunStatusCanceled                    JobRunStatus = "Canceled"
	JobRunStatusCanceling                   JobRunStatus = "Canceling"
	JobRunStatusFailed                      JobRunStatus = "Failed"
	JobRunStatusPausedByBandwidthManagement JobRunStatus = "PausedByBandwidthManagement"
	JobRunStatusQueued                      JobRunStatus = "Queued"
	JobRunStatusRunning                     JobRunStatus = "Running"
	JobRunStatusStarted                     JobRunStatus = "Started"
	JobRunStatusSucceeded                   JobRunStatus = "Succeeded"
)

func PossibleValuesForJobRunStatus() []string {
	return []string{
		string(JobRunStatusCancelRequested),
		string(JobRunStatusCanceled),
		string(JobRunStatusCanceling),
		string(JobRunStatusFailed),
		string(JobRunStatusPausedByBandwidthManagement),
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
	out, err := parseJobRunStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseJobRunStatus(input string) (*JobRunStatus, error) {
	vals := map[string]JobRunStatus{
		"cancelrequested":             JobRunStatusCancelRequested,
		"canceled":                    JobRunStatusCanceled,
		"canceling":                   JobRunStatusCanceling,
		"failed":                      JobRunStatusFailed,
		"pausedbybandwidthmanagement": JobRunStatusPausedByBandwidthManagement,
		"queued":                      JobRunStatusQueued,
		"running":                     JobRunStatusRunning,
		"started":                     JobRunStatusStarted,
		"succeeded":                   JobRunStatusSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JobRunStatus(input)
	return &out, nil
}

type JobType string

const (
	JobTypeCloudToCloud           JobType = "CloudToCloud"
	JobTypeOnPremToCloud          JobType = "OnPremToCloud"
	JobTypeOnPremToCloudAgentLess JobType = "OnPremToCloudAgentLess"
)

func PossibleValuesForJobType() []string {
	return []string{
		string(JobTypeCloudToCloud),
		string(JobTypeOnPremToCloud),
		string(JobTypeOnPremToCloudAgentLess),
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
		"cloudtocloud":           JobTypeCloudToCloud,
		"onpremtocloud":          JobTypeOnPremToCloud,
		"onpremtocloudagentless": JobTypeOnPremToCloudAgentLess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JobType(input)
	return &out, nil
}

type Minute float64

const (
	MinuteThreeZero Minute = 30
	MinuteZero      Minute = 0
)

func PossibleValuesForMinute() []float64 {
	return []float64{
		float64(MinuteThreeZero),
		float64(MinuteZero),
	}
}

type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
	}
}

func (s *ProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"canceled":  ProvisioningStateCanceled,
		"deleting":  ProvisioningStateDeleting,
		"failed":    ProvisioningStateFailed,
		"succeeded": ProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}
