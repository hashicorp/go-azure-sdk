package tasks

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudTask struct {
	AffinityInfo                 *AffinityInformation           `json:"affinityInfo,omitempty"`
	ApplicationPackageReferences *[]ApplicationPackageReference `json:"applicationPackageReferences,omitempty"`
	AuthenticationTokenSettings  *AuthenticationTokenSettings   `json:"authenticationTokenSettings,omitempty"`
	CommandLine                  *string                        `json:"commandLine,omitempty"`
	Constraints                  *TaskConstraints               `json:"constraints,omitempty"`
	ContainerSettings            *TaskContainerSettings         `json:"containerSettings,omitempty"`
	CreationTime                 *string                        `json:"creationTime,omitempty"`
	DependsOn                    *TaskDependencies              `json:"dependsOn,omitempty"`
	DisplayName                  *string                        `json:"displayName,omitempty"`
	ETag                         *string                        `json:"eTag,omitempty"`
	EnvironmentSettings          *[]EnvironmentSetting          `json:"environmentSettings,omitempty"`
	ExecutionInfo                *TaskExecutionInformation      `json:"executionInfo,omitempty"`
	ExitConditions               *ExitConditions                `json:"exitConditions,omitempty"`
	Id                           *string                        `json:"id,omitempty"`
	LastModified                 *string                        `json:"lastModified,omitempty"`
	MultiInstanceSettings        *MultiInstanceSettings         `json:"multiInstanceSettings,omitempty"`
	NodeInfo                     *ComputeNodeInformation        `json:"nodeInfo,omitempty"`
	OutputFiles                  *[]OutputFile                  `json:"outputFiles,omitempty"`
	PreviousState                *TaskState                     `json:"previousState,omitempty"`
	PreviousStateTransitionTime  *string                        `json:"previousStateTransitionTime,omitempty"`
	RequiredSlots                *int64                         `json:"requiredSlots,omitempty"`
	ResourceFiles                *[]ResourceFile                `json:"resourceFiles,omitempty"`
	State                        *TaskState                     `json:"state,omitempty"`
	StateTransitionTime          *string                        `json:"stateTransitionTime,omitempty"`
	Stats                        *TaskStatistics                `json:"stats,omitempty"`
	Url                          *string                        `json:"url,omitempty"`
	UserIdentity                 *UserIdentity                  `json:"userIdentity,omitempty"`
}

func (o *CloudTask) GetCreationTimeAsTime() (*time.Time, error) {
	if o.CreationTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreationTime, "2006-01-02T15:04:05Z07:00")
}

func (o *CloudTask) SetCreationTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreationTime = &formatted
}

func (o *CloudTask) GetLastModifiedAsTime() (*time.Time, error) {
	if o.LastModified == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastModified, "2006-01-02T15:04:05Z07:00")
}

func (o *CloudTask) SetLastModifiedAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastModified = &formatted
}

func (o *CloudTask) GetPreviousStateTransitionTimeAsTime() (*time.Time, error) {
	if o.PreviousStateTransitionTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.PreviousStateTransitionTime, "2006-01-02T15:04:05Z07:00")
}

func (o *CloudTask) SetPreviousStateTransitionTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.PreviousStateTransitionTime = &formatted
}

func (o *CloudTask) GetStateTransitionTimeAsTime() (*time.Time, error) {
	if o.StateTransitionTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StateTransitionTime, "2006-01-02T15:04:05Z07:00")
}

func (o *CloudTask) SetStateTransitionTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StateTransitionTime = &formatted
}
