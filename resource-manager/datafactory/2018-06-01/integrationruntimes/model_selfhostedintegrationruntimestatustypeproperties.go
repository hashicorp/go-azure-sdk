package integrationruntimes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SelfHostedIntegrationRuntimeStatusTypeProperties struct {
	AutoUpdate                               *IntegrationRuntimeAutoUpdate                    `json:"autoUpdate,omitempty"`
	AutoUpdateETA                            *string                                          `json:"autoUpdateETA,omitempty"`
	Capabilities                             *map[string]string                               `json:"capabilities,omitempty"`
	CreateTime                               *string                                          `json:"createTime,omitempty"`
	InternalChannelEncryption                *IntegrationRuntimeInternalChannelEncryptionMode `json:"internalChannelEncryption,omitempty"`
	LatestVersion                            *string                                          `json:"latestVersion,omitempty"`
	Links                                    *[]LinkedIntegrationRuntime                      `json:"links,omitempty"`
	LocalTimeZoneOffset                      *string                                          `json:"localTimeZoneOffset,omitempty"`
	Nodes                                    *[]SelfHostedIntegrationRuntimeNode              `json:"nodes,omitempty"`
	PushedVersion                            *string                                          `json:"pushedVersion,omitempty"`
	ScheduledUpdateDate                      *string                                          `json:"scheduledUpdateDate,omitempty"`
	SelfContainedInteractiveAuthoringEnabled *bool                                            `json:"selfContainedInteractiveAuthoringEnabled,omitempty"`
	ServiceUrls                              *[]string                                        `json:"serviceUrls,omitempty"`
	TaskQueueId                              *string                                          `json:"taskQueueId,omitempty"`
	UpdateDelayOffset                        *string                                          `json:"updateDelayOffset,omitempty"`
	Version                                  *string                                          `json:"version,omitempty"`
	VersionStatus                            *string                                          `json:"versionStatus,omitempty"`
}
