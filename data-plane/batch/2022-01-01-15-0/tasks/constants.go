package tasks

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessScope string

const (
	AccessScopeJob AccessScope = "job"
)

func PossibleValuesForAccessScope() []string {
	return []string{
		string(AccessScopeJob),
	}
}

func (s *AccessScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessScope(input string) (*AccessScope, error) {
	vals := map[string]AccessScope{
		"job": AccessScopeJob,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessScope(input)
	return &out, nil
}

type AutoUserScope string

const (
	AutoUserScopePool AutoUserScope = "pool"
	AutoUserScopeTask AutoUserScope = "task"
)

func PossibleValuesForAutoUserScope() []string {
	return []string{
		string(AutoUserScopePool),
		string(AutoUserScopeTask),
	}
}

func (s *AutoUserScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAutoUserScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAutoUserScope(input string) (*AutoUserScope, error) {
	vals := map[string]AutoUserScope{
		"pool": AutoUserScopePool,
		"task": AutoUserScopeTask,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutoUserScope(input)
	return &out, nil
}

type ContainerWorkingDirectory string

const (
	ContainerWorkingDirectoryContainerImageDefault ContainerWorkingDirectory = "containerImageDefault"
	ContainerWorkingDirectoryTaskWorkingDirectory  ContainerWorkingDirectory = "taskWorkingDirectory"
)

func PossibleValuesForContainerWorkingDirectory() []string {
	return []string{
		string(ContainerWorkingDirectoryContainerImageDefault),
		string(ContainerWorkingDirectoryTaskWorkingDirectory),
	}
}

func (s *ContainerWorkingDirectory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseContainerWorkingDirectory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseContainerWorkingDirectory(input string) (*ContainerWorkingDirectory, error) {
	vals := map[string]ContainerWorkingDirectory{
		"containerimagedefault": ContainerWorkingDirectoryContainerImageDefault,
		"taskworkingdirectory":  ContainerWorkingDirectoryTaskWorkingDirectory,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContainerWorkingDirectory(input)
	return &out, nil
}

type DependencyAction string

const (
	DependencyActionBlock   DependencyAction = "block"
	DependencyActionSatisfy DependencyAction = "satisfy"
)

func PossibleValuesForDependencyAction() []string {
	return []string{
		string(DependencyActionBlock),
		string(DependencyActionSatisfy),
	}
}

func (s *DependencyAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDependencyAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDependencyAction(input string) (*DependencyAction, error) {
	vals := map[string]DependencyAction{
		"block":   DependencyActionBlock,
		"satisfy": DependencyActionSatisfy,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DependencyAction(input)
	return &out, nil
}

type ElevationLevel string

const (
	ElevationLevelAdmin    ElevationLevel = "admin"
	ElevationLevelNonadmin ElevationLevel = "nonadmin"
)

func PossibleValuesForElevationLevel() []string {
	return []string{
		string(ElevationLevelAdmin),
		string(ElevationLevelNonadmin),
	}
}

func (s *ElevationLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseElevationLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseElevationLevel(input string) (*ElevationLevel, error) {
	vals := map[string]ElevationLevel{
		"admin":    ElevationLevelAdmin,
		"nonadmin": ElevationLevelNonadmin,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ElevationLevel(input)
	return &out, nil
}

type ErrorCategory string

const (
	ErrorCategoryServererror ErrorCategory = "servererror"
	ErrorCategoryUsererror   ErrorCategory = "usererror"
)

func PossibleValuesForErrorCategory() []string {
	return []string{
		string(ErrorCategoryServererror),
		string(ErrorCategoryUsererror),
	}
}

func (s *ErrorCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseErrorCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseErrorCategory(input string) (*ErrorCategory, error) {
	vals := map[string]ErrorCategory{
		"servererror": ErrorCategoryServererror,
		"usererror":   ErrorCategoryUsererror,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ErrorCategory(input)
	return &out, nil
}

type JobAction string

const (
	JobActionDisable   JobAction = "disable"
	JobActionNone      JobAction = "none"
	JobActionTerminate JobAction = "terminate"
)

func PossibleValuesForJobAction() []string {
	return []string{
		string(JobActionDisable),
		string(JobActionNone),
		string(JobActionTerminate),
	}
}

func (s *JobAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseJobAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseJobAction(input string) (*JobAction, error) {
	vals := map[string]JobAction{
		"disable":   JobActionDisable,
		"none":      JobActionNone,
		"terminate": JobActionTerminate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JobAction(input)
	return &out, nil
}

type OutputFileUploadCondition string

const (
	OutputFileUploadConditionTaskcompletion OutputFileUploadCondition = "taskcompletion"
	OutputFileUploadConditionTaskfailure    OutputFileUploadCondition = "taskfailure"
	OutputFileUploadConditionTasksuccess    OutputFileUploadCondition = "tasksuccess"
)

func PossibleValuesForOutputFileUploadCondition() []string {
	return []string{
		string(OutputFileUploadConditionTaskcompletion),
		string(OutputFileUploadConditionTaskfailure),
		string(OutputFileUploadConditionTasksuccess),
	}
}

func (s *OutputFileUploadCondition) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOutputFileUploadCondition(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOutputFileUploadCondition(input string) (*OutputFileUploadCondition, error) {
	vals := map[string]OutputFileUploadCondition{
		"taskcompletion": OutputFileUploadConditionTaskcompletion,
		"taskfailure":    OutputFileUploadConditionTaskfailure,
		"tasksuccess":    OutputFileUploadConditionTasksuccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OutputFileUploadCondition(input)
	return &out, nil
}

type SubtaskState string

const (
	SubtaskStateCompleted SubtaskState = "completed"
	SubtaskStatePreparing SubtaskState = "preparing"
	SubtaskStateRunning   SubtaskState = "running"
)

func PossibleValuesForSubtaskState() []string {
	return []string{
		string(SubtaskStateCompleted),
		string(SubtaskStatePreparing),
		string(SubtaskStateRunning),
	}
}

func (s *SubtaskState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSubtaskState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSubtaskState(input string) (*SubtaskState, error) {
	vals := map[string]SubtaskState{
		"completed": SubtaskStateCompleted,
		"preparing": SubtaskStatePreparing,
		"running":   SubtaskStateRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubtaskState(input)
	return &out, nil
}

type TaskAddStatus string

const (
	TaskAddStatusClienterror TaskAddStatus = "clienterror"
	TaskAddStatusServererror TaskAddStatus = "servererror"
	TaskAddStatusSuccess     TaskAddStatus = "success"
)

func PossibleValuesForTaskAddStatus() []string {
	return []string{
		string(TaskAddStatusClienterror),
		string(TaskAddStatusServererror),
		string(TaskAddStatusSuccess),
	}
}

func (s *TaskAddStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTaskAddStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTaskAddStatus(input string) (*TaskAddStatus, error) {
	vals := map[string]TaskAddStatus{
		"clienterror": TaskAddStatusClienterror,
		"servererror": TaskAddStatusServererror,
		"success":     TaskAddStatusSuccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TaskAddStatus(input)
	return &out, nil
}

type TaskExecutionResult string

const (
	TaskExecutionResultFailure TaskExecutionResult = "failure"
	TaskExecutionResultSuccess TaskExecutionResult = "success"
)

func PossibleValuesForTaskExecutionResult() []string {
	return []string{
		string(TaskExecutionResultFailure),
		string(TaskExecutionResultSuccess),
	}
}

func (s *TaskExecutionResult) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTaskExecutionResult(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTaskExecutionResult(input string) (*TaskExecutionResult, error) {
	vals := map[string]TaskExecutionResult{
		"failure": TaskExecutionResultFailure,
		"success": TaskExecutionResultSuccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TaskExecutionResult(input)
	return &out, nil
}

type TaskState string

const (
	TaskStateActive    TaskState = "active"
	TaskStateCompleted TaskState = "completed"
	TaskStatePreparing TaskState = "preparing"
	TaskStateRunning   TaskState = "running"
)

func PossibleValuesForTaskState() []string {
	return []string{
		string(TaskStateActive),
		string(TaskStateCompleted),
		string(TaskStatePreparing),
		string(TaskStateRunning),
	}
}

func (s *TaskState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTaskState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTaskState(input string) (*TaskState, error) {
	vals := map[string]TaskState{
		"active":    TaskStateActive,
		"completed": TaskStateCompleted,
		"preparing": TaskStatePreparing,
		"running":   TaskStateRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TaskState(input)
	return &out, nil
}
