package shares

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ShareProperties struct {
	AccessProtocol     ShareAccessProtocol  `json:"accessProtocol"`
	AzureContainerInfo *AzureContainerInfo  `json:"azureContainerInfo,omitempty"`
	ClientAccessRights *[]ClientAccessRight `json:"clientAccessRights,omitempty"`
	DataPolicy         *DataPolicy          `json:"dataPolicy,omitempty"`
	Description        *string              `json:"description,omitempty"`
	MonitoringStatus   MonitoringStatus     `json:"monitoringStatus"`
	RefreshDetails     *RefreshDetails      `json:"refreshDetails,omitempty"`
	ShareMappings      *[]MountPointMap     `json:"shareMappings,omitempty"`
	ShareStatus        ShareStatus          `json:"shareStatus"`
	UserAccessRights   *[]UserAccessRight   `json:"userAccessRights,omitempty"`
}
