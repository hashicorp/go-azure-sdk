package computenodes

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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

type CertificateStoreLocation string

const (
	CertificateStoreLocationCurrentuser  CertificateStoreLocation = "currentuser"
	CertificateStoreLocationLocalmachine CertificateStoreLocation = "localmachine"
)

func PossibleValuesForCertificateStoreLocation() []string {
	return []string{
		string(CertificateStoreLocationCurrentuser),
		string(CertificateStoreLocationLocalmachine),
	}
}

func (s *CertificateStoreLocation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCertificateStoreLocation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCertificateStoreLocation(input string) (*CertificateStoreLocation, error) {
	vals := map[string]CertificateStoreLocation{
		"currentuser":  CertificateStoreLocationCurrentuser,
		"localmachine": CertificateStoreLocationLocalmachine,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CertificateStoreLocation(input)
	return &out, nil
}

type CertificateVisibility string

const (
	CertificateVisibilityRemoteuser CertificateVisibility = "remoteuser"
	CertificateVisibilityStarttask  CertificateVisibility = "starttask"
	CertificateVisibilityTask       CertificateVisibility = "task"
)

func PossibleValuesForCertificateVisibility() []string {
	return []string{
		string(CertificateVisibilityRemoteuser),
		string(CertificateVisibilityStarttask),
		string(CertificateVisibilityTask),
	}
}

func (s *CertificateVisibility) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCertificateVisibility(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCertificateVisibility(input string) (*CertificateVisibility, error) {
	vals := map[string]CertificateVisibility{
		"remoteuser": CertificateVisibilityRemoteuser,
		"starttask":  CertificateVisibilityStarttask,
		"task":       CertificateVisibilityTask,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CertificateVisibility(input)
	return &out, nil
}

type ComputeNodeDeallocationOption string

const (
	ComputeNodeDeallocationOptionRequeue        ComputeNodeDeallocationOption = "requeue"
	ComputeNodeDeallocationOptionRetaineddata   ComputeNodeDeallocationOption = "retaineddata"
	ComputeNodeDeallocationOptionTaskcompletion ComputeNodeDeallocationOption = "taskcompletion"
	ComputeNodeDeallocationOptionTerminate      ComputeNodeDeallocationOption = "terminate"
)

func PossibleValuesForComputeNodeDeallocationOption() []string {
	return []string{
		string(ComputeNodeDeallocationOptionRequeue),
		string(ComputeNodeDeallocationOptionRetaineddata),
		string(ComputeNodeDeallocationOptionTaskcompletion),
		string(ComputeNodeDeallocationOptionTerminate),
	}
}

func (s *ComputeNodeDeallocationOption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseComputeNodeDeallocationOption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseComputeNodeDeallocationOption(input string) (*ComputeNodeDeallocationOption, error) {
	vals := map[string]ComputeNodeDeallocationOption{
		"requeue":        ComputeNodeDeallocationOptionRequeue,
		"retaineddata":   ComputeNodeDeallocationOptionRetaineddata,
		"taskcompletion": ComputeNodeDeallocationOptionTaskcompletion,
		"terminate":      ComputeNodeDeallocationOptionTerminate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComputeNodeDeallocationOption(input)
	return &out, nil
}

type ComputeNodeRebootOption string

const (
	ComputeNodeRebootOptionRequeue        ComputeNodeRebootOption = "requeue"
	ComputeNodeRebootOptionRetaineddata   ComputeNodeRebootOption = "retaineddata"
	ComputeNodeRebootOptionTaskcompletion ComputeNodeRebootOption = "taskcompletion"
	ComputeNodeRebootOptionTerminate      ComputeNodeRebootOption = "terminate"
)

func PossibleValuesForComputeNodeRebootOption() []string {
	return []string{
		string(ComputeNodeRebootOptionRequeue),
		string(ComputeNodeRebootOptionRetaineddata),
		string(ComputeNodeRebootOptionTaskcompletion),
		string(ComputeNodeRebootOptionTerminate),
	}
}

func (s *ComputeNodeRebootOption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseComputeNodeRebootOption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseComputeNodeRebootOption(input string) (*ComputeNodeRebootOption, error) {
	vals := map[string]ComputeNodeRebootOption{
		"requeue":        ComputeNodeRebootOptionRequeue,
		"retaineddata":   ComputeNodeRebootOptionRetaineddata,
		"taskcompletion": ComputeNodeRebootOptionTaskcompletion,
		"terminate":      ComputeNodeRebootOptionTerminate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComputeNodeRebootOption(input)
	return &out, nil
}

type ComputeNodeReimageOption string

const (
	ComputeNodeReimageOptionRequeue        ComputeNodeReimageOption = "requeue"
	ComputeNodeReimageOptionRetaineddata   ComputeNodeReimageOption = "retaineddata"
	ComputeNodeReimageOptionTaskcompletion ComputeNodeReimageOption = "taskcompletion"
	ComputeNodeReimageOptionTerminate      ComputeNodeReimageOption = "terminate"
)

func PossibleValuesForComputeNodeReimageOption() []string {
	return []string{
		string(ComputeNodeReimageOptionRequeue),
		string(ComputeNodeReimageOptionRetaineddata),
		string(ComputeNodeReimageOptionTaskcompletion),
		string(ComputeNodeReimageOptionTerminate),
	}
}

func (s *ComputeNodeReimageOption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseComputeNodeReimageOption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseComputeNodeReimageOption(input string) (*ComputeNodeReimageOption, error) {
	vals := map[string]ComputeNodeReimageOption{
		"requeue":        ComputeNodeReimageOptionRequeue,
		"retaineddata":   ComputeNodeReimageOptionRetaineddata,
		"taskcompletion": ComputeNodeReimageOptionTaskcompletion,
		"terminate":      ComputeNodeReimageOptionTerminate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComputeNodeReimageOption(input)
	return &out, nil
}

type ComputeNodeState string

const (
	ComputeNodeStateCreating            ComputeNodeState = "creating"
	ComputeNodeStateIdle                ComputeNodeState = "idle"
	ComputeNodeStateLeavingpool         ComputeNodeState = "leavingpool"
	ComputeNodeStateOffline             ComputeNodeState = "offline"
	ComputeNodeStatePreempted           ComputeNodeState = "preempted"
	ComputeNodeStateRebooting           ComputeNodeState = "rebooting"
	ComputeNodeStateReimaging           ComputeNodeState = "reimaging"
	ComputeNodeStateRunning             ComputeNodeState = "running"
	ComputeNodeStateStarting            ComputeNodeState = "starting"
	ComputeNodeStateStarttaskfailed     ComputeNodeState = "starttaskfailed"
	ComputeNodeStateUnknown             ComputeNodeState = "unknown"
	ComputeNodeStateUnusable            ComputeNodeState = "unusable"
	ComputeNodeStateWaitingforstarttask ComputeNodeState = "waitingforstarttask"
)

func PossibleValuesForComputeNodeState() []string {
	return []string{
		string(ComputeNodeStateCreating),
		string(ComputeNodeStateIdle),
		string(ComputeNodeStateLeavingpool),
		string(ComputeNodeStateOffline),
		string(ComputeNodeStatePreempted),
		string(ComputeNodeStateRebooting),
		string(ComputeNodeStateReimaging),
		string(ComputeNodeStateRunning),
		string(ComputeNodeStateStarting),
		string(ComputeNodeStateStarttaskfailed),
		string(ComputeNodeStateUnknown),
		string(ComputeNodeStateUnusable),
		string(ComputeNodeStateWaitingforstarttask),
	}
}

func (s *ComputeNodeState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseComputeNodeState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseComputeNodeState(input string) (*ComputeNodeState, error) {
	vals := map[string]ComputeNodeState{
		"creating":            ComputeNodeStateCreating,
		"idle":                ComputeNodeStateIdle,
		"leavingpool":         ComputeNodeStateLeavingpool,
		"offline":             ComputeNodeStateOffline,
		"preempted":           ComputeNodeStatePreempted,
		"rebooting":           ComputeNodeStateRebooting,
		"reimaging":           ComputeNodeStateReimaging,
		"running":             ComputeNodeStateRunning,
		"starting":            ComputeNodeStateStarting,
		"starttaskfailed":     ComputeNodeStateStarttaskfailed,
		"unknown":             ComputeNodeStateUnknown,
		"unusable":            ComputeNodeStateUnusable,
		"waitingforstarttask": ComputeNodeStateWaitingforstarttask,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComputeNodeState(input)
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

type DisableComputeNodeSchedulingOption string

const (
	DisableComputeNodeSchedulingOptionRequeue        DisableComputeNodeSchedulingOption = "requeue"
	DisableComputeNodeSchedulingOptionTaskcompletion DisableComputeNodeSchedulingOption = "taskcompletion"
	DisableComputeNodeSchedulingOptionTerminate      DisableComputeNodeSchedulingOption = "terminate"
)

func PossibleValuesForDisableComputeNodeSchedulingOption() []string {
	return []string{
		string(DisableComputeNodeSchedulingOptionRequeue),
		string(DisableComputeNodeSchedulingOptionTaskcompletion),
		string(DisableComputeNodeSchedulingOptionTerminate),
	}
}

func (s *DisableComputeNodeSchedulingOption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDisableComputeNodeSchedulingOption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDisableComputeNodeSchedulingOption(input string) (*DisableComputeNodeSchedulingOption, error) {
	vals := map[string]DisableComputeNodeSchedulingOption{
		"requeue":        DisableComputeNodeSchedulingOptionRequeue,
		"taskcompletion": DisableComputeNodeSchedulingOptionTaskcompletion,
		"terminate":      DisableComputeNodeSchedulingOptionTerminate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DisableComputeNodeSchedulingOption(input)
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

type InboundEndpointProtocol string

const (
	InboundEndpointProtocolTcp InboundEndpointProtocol = "tcp"
	InboundEndpointProtocolUdp InboundEndpointProtocol = "udp"
)

func PossibleValuesForInboundEndpointProtocol() []string {
	return []string{
		string(InboundEndpointProtocolTcp),
		string(InboundEndpointProtocolUdp),
	}
}

func (s *InboundEndpointProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInboundEndpointProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInboundEndpointProtocol(input string) (*InboundEndpointProtocol, error) {
	vals := map[string]InboundEndpointProtocol{
		"tcp": InboundEndpointProtocolTcp,
		"udp": InboundEndpointProtocolUdp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InboundEndpointProtocol(input)
	return &out, nil
}

type SchedulingState string

const (
	SchedulingStateDisabled SchedulingState = "disabled"
	SchedulingStateEnabled  SchedulingState = "enabled"
)

func PossibleValuesForSchedulingState() []string {
	return []string{
		string(SchedulingStateDisabled),
		string(SchedulingStateEnabled),
	}
}

func (s *SchedulingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSchedulingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSchedulingState(input string) (*SchedulingState, error) {
	vals := map[string]SchedulingState{
		"disabled": SchedulingStateDisabled,
		"enabled":  SchedulingStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SchedulingState(input)
	return &out, nil
}

type StartTaskState string

const (
	StartTaskStateCompleted StartTaskState = "completed"
	StartTaskStateRunning   StartTaskState = "running"
)

func PossibleValuesForStartTaskState() []string {
	return []string{
		string(StartTaskStateCompleted),
		string(StartTaskStateRunning),
	}
}

func (s *StartTaskState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStartTaskState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStartTaskState(input string) (*StartTaskState, error) {
	vals := map[string]StartTaskState{
		"completed": StartTaskStateCompleted,
		"running":   StartTaskStateRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StartTaskState(input)
	return &out, nil
}

type StatusLevelTypes string

const (
	StatusLevelTypesError   StatusLevelTypes = "Error"
	StatusLevelTypesInfo    StatusLevelTypes = "Info"
	StatusLevelTypesWarning StatusLevelTypes = "Warning"
)

func PossibleValuesForStatusLevelTypes() []string {
	return []string{
		string(StatusLevelTypesError),
		string(StatusLevelTypesInfo),
		string(StatusLevelTypesWarning),
	}
}

func (s *StatusLevelTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStatusLevelTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStatusLevelTypes(input string) (*StatusLevelTypes, error) {
	vals := map[string]StatusLevelTypes{
		"error":   StatusLevelTypesError,
		"info":    StatusLevelTypesInfo,
		"warning": StatusLevelTypesWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StatusLevelTypes(input)
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
