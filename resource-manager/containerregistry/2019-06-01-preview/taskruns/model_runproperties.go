package taskruns

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RunProperties struct {
	AgentConfiguration *AgentProperties         `json:"agentConfiguration,omitempty"`
	AgentPoolName      *string                  `json:"agentPoolName,omitempty"`
	CreateTime         *string                  `json:"createTime,omitempty"`
	CustomRegistries   *[]string                `json:"customRegistries,omitempty"`
	FinishTime         *string                  `json:"finishTime,omitempty"`
	ImageUpdateTrigger *ImageUpdateTrigger      `json:"imageUpdateTrigger,omitempty"`
	IsArchiveEnabled   *bool                    `json:"isArchiveEnabled,omitempty"`
	LastUpdatedTime    *string                  `json:"lastUpdatedTime,omitempty"`
	LogArtifact        *ImageDescriptor         `json:"logArtifact,omitempty"`
	OutputImages       *[]ImageDescriptor       `json:"outputImages,omitempty"`
	Platform           *PlatformProperties      `json:"platform,omitempty"`
	ProvisioningState  *ProvisioningState       `json:"provisioningState,omitempty"`
	RunErrorMessage    *string                  `json:"runErrorMessage,omitempty"`
	RunId              *string                  `json:"runId,omitempty"`
	RunType            *RunType                 `json:"runType,omitempty"`
	SourceRegistryAuth *string                  `json:"sourceRegistryAuth,omitempty"`
	SourceTrigger      *SourceTriggerDescriptor `json:"sourceTrigger,omitempty"`
	StartTime          *string                  `json:"startTime,omitempty"`
	Status             *RunStatus               `json:"status,omitempty"`
	Task               *string                  `json:"task,omitempty"`
	TimerTrigger       *TimerTriggerDescriptor  `json:"timerTrigger,omitempty"`
	UpdateTriggerToken *string                  `json:"updateTriggerToken,omitempty"`
}
