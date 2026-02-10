package tasks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TaskAddParameter struct {
	AffinityInfo                 *AffinityInformation           `json:"affinityInfo,omitempty"`
	ApplicationPackageReferences *[]ApplicationPackageReference `json:"applicationPackageReferences,omitempty"`
	AuthenticationTokenSettings  *AuthenticationTokenSettings   `json:"authenticationTokenSettings,omitempty"`
	CommandLine                  string                         `json:"commandLine"`
	Constraints                  *TaskConstraints               `json:"constraints,omitempty"`
	ContainerSettings            *TaskContainerSettings         `json:"containerSettings,omitempty"`
	DependsOn                    *TaskDependencies              `json:"dependsOn,omitempty"`
	DisplayName                  *string                        `json:"displayName,omitempty"`
	EnvironmentSettings          *[]EnvironmentSetting          `json:"environmentSettings,omitempty"`
	ExitConditions               *ExitConditions                `json:"exitConditions,omitempty"`
	Id                           string                         `json:"id"`
	MultiInstanceSettings        *MultiInstanceSettings         `json:"multiInstanceSettings,omitempty"`
	OutputFiles                  *[]OutputFile                  `json:"outputFiles,omitempty"`
	RequiredSlots                *int64                         `json:"requiredSlots,omitempty"`
	ResourceFiles                *[]ResourceFile                `json:"resourceFiles,omitempty"`
	UserIdentity                 *UserIdentity                  `json:"userIdentity,omitempty"`
}
