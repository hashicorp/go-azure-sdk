package sparkjobdefinitions

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BigDataPoolReferenceType string

const (
	BigDataPoolReferenceTypeBigDataPoolReference BigDataPoolReferenceType = "BigDataPoolReference"
)

func PossibleValuesForBigDataPoolReferenceType() []string {
	return []string{
		string(BigDataPoolReferenceTypeBigDataPoolReference),
	}
}

func (s *BigDataPoolReferenceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBigDataPoolReferenceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBigDataPoolReferenceType(input string) (*BigDataPoolReferenceType, error) {
	vals := map[string]BigDataPoolReferenceType{
		"bigdatapoolreference": BigDataPoolReferenceTypeBigDataPoolReference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BigDataPoolReferenceType(input)
	return &out, nil
}

type PluginCurrentState string

const (
	PluginCurrentStateCleanup             PluginCurrentState = "Cleanup"
	PluginCurrentStateEnded               PluginCurrentState = "Ended"
	PluginCurrentStateMonitoring          PluginCurrentState = "Monitoring"
	PluginCurrentStatePreparation         PluginCurrentState = "Preparation"
	PluginCurrentStateQueued              PluginCurrentState = "Queued"
	PluginCurrentStateResourceAcquisition PluginCurrentState = "ResourceAcquisition"
	PluginCurrentStateSubmission          PluginCurrentState = "Submission"
)

func PossibleValuesForPluginCurrentState() []string {
	return []string{
		string(PluginCurrentStateCleanup),
		string(PluginCurrentStateEnded),
		string(PluginCurrentStateMonitoring),
		string(PluginCurrentStatePreparation),
		string(PluginCurrentStateQueued),
		string(PluginCurrentStateResourceAcquisition),
		string(PluginCurrentStateSubmission),
	}
}

func (s *PluginCurrentState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePluginCurrentState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePluginCurrentState(input string) (*PluginCurrentState, error) {
	vals := map[string]PluginCurrentState{
		"cleanup":             PluginCurrentStateCleanup,
		"ended":               PluginCurrentStateEnded,
		"monitoring":          PluginCurrentStateMonitoring,
		"preparation":         PluginCurrentStatePreparation,
		"queued":              PluginCurrentStateQueued,
		"resourceacquisition": PluginCurrentStateResourceAcquisition,
		"submission":          PluginCurrentStateSubmission,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PluginCurrentState(input)
	return &out, nil
}

type SchedulerCurrentState string

const (
	SchedulerCurrentStateEnded     SchedulerCurrentState = "Ended"
	SchedulerCurrentStateQueued    SchedulerCurrentState = "Queued"
	SchedulerCurrentStateScheduled SchedulerCurrentState = "Scheduled"
)

func PossibleValuesForSchedulerCurrentState() []string {
	return []string{
		string(SchedulerCurrentStateEnded),
		string(SchedulerCurrentStateQueued),
		string(SchedulerCurrentStateScheduled),
	}
}

func (s *SchedulerCurrentState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSchedulerCurrentState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSchedulerCurrentState(input string) (*SchedulerCurrentState, error) {
	vals := map[string]SchedulerCurrentState{
		"ended":     SchedulerCurrentStateEnded,
		"queued":    SchedulerCurrentStateQueued,
		"scheduled": SchedulerCurrentStateScheduled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SchedulerCurrentState(input)
	return &out, nil
}

type SparkBatchJobResultType string

const (
	SparkBatchJobResultTypeCancelled SparkBatchJobResultType = "Cancelled"
	SparkBatchJobResultTypeFailed    SparkBatchJobResultType = "Failed"
	SparkBatchJobResultTypeSucceeded SparkBatchJobResultType = "Succeeded"
	SparkBatchJobResultTypeUncertain SparkBatchJobResultType = "Uncertain"
)

func PossibleValuesForSparkBatchJobResultType() []string {
	return []string{
		string(SparkBatchJobResultTypeCancelled),
		string(SparkBatchJobResultTypeFailed),
		string(SparkBatchJobResultTypeSucceeded),
		string(SparkBatchJobResultTypeUncertain),
	}
}

func (s *SparkBatchJobResultType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSparkBatchJobResultType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSparkBatchJobResultType(input string) (*SparkBatchJobResultType, error) {
	vals := map[string]SparkBatchJobResultType{
		"cancelled": SparkBatchJobResultTypeCancelled,
		"failed":    SparkBatchJobResultTypeFailed,
		"succeeded": SparkBatchJobResultTypeSucceeded,
		"uncertain": SparkBatchJobResultTypeUncertain,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SparkBatchJobResultType(input)
	return &out, nil
}

type SparkErrorSource string

const (
	SparkErrorSourceDependency SparkErrorSource = "Dependency"
	SparkErrorSourceSystem     SparkErrorSource = "System"
	SparkErrorSourceUnknown    SparkErrorSource = "Unknown"
	SparkErrorSourceUser       SparkErrorSource = "User"
)

func PossibleValuesForSparkErrorSource() []string {
	return []string{
		string(SparkErrorSourceDependency),
		string(SparkErrorSourceSystem),
		string(SparkErrorSourceUnknown),
		string(SparkErrorSourceUser),
	}
}

func (s *SparkErrorSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSparkErrorSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSparkErrorSource(input string) (*SparkErrorSource, error) {
	vals := map[string]SparkErrorSource{
		"dependency": SparkErrorSourceDependency,
		"system":     SparkErrorSourceSystem,
		"unknown":    SparkErrorSourceUnknown,
		"user":       SparkErrorSourceUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SparkErrorSource(input)
	return &out, nil
}

type SparkJobType string

const (
	SparkJobTypeSparkBatch   SparkJobType = "SparkBatch"
	SparkJobTypeSparkSession SparkJobType = "SparkSession"
)

func PossibleValuesForSparkJobType() []string {
	return []string{
		string(SparkJobTypeSparkBatch),
		string(SparkJobTypeSparkSession),
	}
}

func (s *SparkJobType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSparkJobType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSparkJobType(input string) (*SparkJobType, error) {
	vals := map[string]SparkJobType{
		"sparkbatch":   SparkJobTypeSparkBatch,
		"sparksession": SparkJobTypeSparkSession,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SparkJobType(input)
	return &out, nil
}
